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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// AddFeature adds a new feature to an existing product on the database together with flags for all environments
// od the product with default values. Returns ID if successful, empty string and error otherwise
func (pr ProductRepository) AddFeature(product domain.Product, feat domain.Feature, envFlags []domain.EnvironmentFlag) (string, error) {
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
		idDoc := bson.D{{"_id", productDAO.ID}}
		upDoc := bson.D{{"$push", bson.M{"features": newFeat}}}
		var updateOpts options.UpdateOptions
		updateOpts.SetUpsert(false)

		if _, err := productCollection.UpdateOne(sessCtx, idDoc, upDoc, &updateOpts); err != nil {
			return nil, err
		}

		for _, envFlag := range envFlags {
			idDoc := bson.D{{"environmentID", envFlag.EnvironmentID}}
			upDoc := bson.D{{"$push", bson.M{"flags": envFlag.Flags[0]}}}
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
		return "", err
	}
	log.Info().Msgf("result: %v\n", result)
	return newFeat.Key, nil
}

// UpdateFeature updates an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) UpdateFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}

// DisableFeature disables an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) DisableFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}
