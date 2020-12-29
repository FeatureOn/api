package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/FeatureOn/api/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/adapters/data/mongodb/mappers"
	"github.com/FeatureOn/api/domain"
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
	productDAO := mappers.MapProduct2ProductDAO(product)
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
		idDoc := bson.D{{"_id", productDAO.ID}}
		upDoc := bson.D{{"$push", bson.M{"environments": newEnv}}}
		var updateOpts options.UpdateOptions
		updateOpts.SetUpsert(false)

		if _, err := productCollection.UpdateOne(sessCtx, idDoc, upDoc, &updateOpts); err != nil {
			return nil, err
		}
		environmentFlag.EnvironmentID = newEnvID.Hex()

		if _, err := flagCollection.InsertOne(sessCtx, mappers.MapEnvironmentFlag2EnvironmentFlagDAO(environmentFlag)); err != nil {
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
	return errors.New("Not implemented")
}
