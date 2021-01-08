package mongodb

import (
	"context"
	"errors"
	"time"

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

// AddFeature adds a new feature to an existing product on the database together with flags for all environments
// od the product with default values. Returns ID if successful, empty string and error otherwise
func (pr ProductRepository) AddFeature(product domain.Product, feat domain.Feature, envFlags []domain.EnvironmentFlag) error {
	productDAO := mappers.MapProduct2ProductDAO(product)
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	productCollection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"), wcMajorityCollectionOpts)
	flagCollection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("FlagsCollection"), wcMajorityCollectionOpts)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	newFeat := mappers.MapFeature2FeatureDAO(feat)
	productDAO.Features = append(productDAO.Features, newFeat)

	// Step 1: Define the callback that specifies the sequence of operations to perform inside the transaction.
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		idDoc := bson.D{{Key: "_id", Value: productDAO.ID}}
		upDoc := bson.D{{Key: "$push", Value: bson.M{"features": newFeat}}}
		var updateOpts options.UpdateOptions
		updateOpts.SetUpsert(false)

		if _, err := productCollection.UpdateOne(sessCtx, idDoc, upDoc, &updateOpts); err != nil {
			return nil, err
		}

		for _, envFlag := range envFlags {
			idDoc := bson.D{{Key: "environmentID", Value: envFlag.EnvironmentID}}
			upDoc := bson.D{{Key: "$push", Value: bson.M{"flags": envFlag.Flags[0]}}}
			var updateOpts options.UpdateOptions
			updateOpts.SetUpsert(true)
			if _, err := flagCollection.UpdateOne(sessCtx, idDoc, upDoc, &updateOpts); err != nil {
				return nil, err
			}
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
		log.Error().Err(err).Msgf("Error adding feature with key %s", newFeat.Key)
		return err
	}
	log.Info().Msgf("result: %v\n", result)
	return nil
}

// UpdateFeature updates an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) UpdateFeature(product domain.Product, feat domain.Feature) error {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse productID: %s into ObjectID", product.ID)
		return errors.New("ProductID format is not as expected")
	}

	idDoc := bson.M{"_id": id, "features.key": feat.Key}
	upDoc := bson.D{{Key: "$set", Value: bson.M{"features.$.name": feat.Name, "features.$.description": feat.Description, "features.$.defaultstate": feat.DefaultState}}}
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	_, err = collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
	if err != nil {
		log.Error().Err(err).Msgf("Error updating the product with productID: %s", product.ID)
		return errors.New("Error updating the product")
	}
	return nil
}

// DisableFeature disables an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) DisableFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}
