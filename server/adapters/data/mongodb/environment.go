package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/FeatureOn/api/server/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/server/adapters/data/mongodb/mappers"
	"github.com/FeatureOn/api/server/domain"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// AddEnvironment adds a new environment together with all its flags with default values and returns its ID,
// returns empty string and error otherwise
func (pr ProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	productDAO, err := mappers.MapProduct2ProductDAO(product)
	if err != nil {
		return "", err
	}
	// Preperation for inter-collection transactions
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	productCollection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"), wcMajorityCollectionOpts)
	flagCollection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("FlagsCollection"), wcMajorityCollectionOpts)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	newEnvID := primitive.NewObjectID()
	newEnv := dao.EnvironmentDAO{
		ID:   newEnvID,
		Name: environmentName,
	}

	// Step 1: Define the callback that specifies the sequence of operations to perform inside the transaction.
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		idDoc := bson.D{{Key: "_id", Value: productDAO.ID}}
		upDoc := bson.D{{Key: "$push", Value: bson.M{"environments": newEnv}}}
		var updateOpts options.UpdateOptions
		updateOpts.SetUpsert(false)

		if _, err := productCollection.UpdateOne(sessCtx, idDoc, upDoc, &updateOpts); err != nil {
			return nil, err
		}
		environmentFlag.EnvironmentID = newEnvID.Hex()

		envFlagDAO, err := mappers.MapEnvironmentFlag2EnvironmentFlagDAO(environmentFlag)
		if err != nil {
			return nil, err
		}
		if _, err := flagCollection.InsertOne(sessCtx, envFlagDAO); err != nil {
			return nil, err
		}

		return nil, nil
	}
	// Step 2: Start a session and run the callback using WithTransaction.
	session, err := pr.dbClient.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		log.Error().Err(err).Msgf("Error adding environment with name %s", environmentName)
		return "", err
	}
	log.Info().Msgf("result: %v\n", result)
	return newEnvID.Hex(), nil
}

// UpdateEnvironment updates an existing environment on the database
func (pr ProductRepository) UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse productID: %s into ObjectID", product.ID)
		return errors.New("ProductID format is not as expected")
	}
	envID, err := primitive.ObjectIDFromHex(environmentID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse environmentID: %s into ObjectID", environmentID)
		return errors.New("EnvronmentID format is not as expected")
	}

	idDoc := bson.M{"_id": id, "environments.id": envID}
	upDoc := bson.D{{Key: "$set", Value: bson.M{"environments.$.name": environmentName}}}
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	_, err = collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
	if err != nil {
		log.Error().Err(err).Msgf("Error updating the product with productID: %s", product.ID)
		return errors.New("Error updating the product")
	}
	return nil
}
