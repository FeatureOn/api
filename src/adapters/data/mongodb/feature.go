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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddFeature adds a new feature to an existing product on the database together with flags for all environments
// od the product with default values. Returns ID if successful, empty string and error otherwise
func (pr ProductRepository) AddFeature(product domain.Product, feat domain.Feature, envFlags []domain.EnvironmentFlag) (string, error) {
	productDAO := mappers.MapProduct2ProductDAO(product)
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	newFeat := mappers.MapFeature2FeatureDAO(feat)
	productDAO.Features = append(productDAO.Features, newFeat)
	idDoc := bson.D{{"_id", productDAO.ID}}
	upDoc := bson.D{{"$set", bson.M{"features": productDAO.Features}}}
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	result, err := collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
	if err == nil {
		if result.MatchedCount == 1 {
			collection = pr.dbClient.Database(pr.dbName).Collection(viper.GetString("FlagsCollection"))
			for _, envFlag := range envFlags {
				idDoc := bson.D{{"environmentID", envFlag.EnvironmentID}}
				upDoc := bson.D{{"$push", bson.M{"flags": envFlag.Flags[0]}}}
				var updateOpts options.UpdateOptions
				updateOpts.SetUpsert(true)
				result, err = collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
			}
			return newFeat.Key, nil
		}
		log.Error().Err(err).Msgf("The productID %s did not match any products in the database", product.ID)
		return "", errors.New("Product not found")
	}
	log.Error().Err(err).Msgf("Error adding feature with key %s", newFeat.Key)
	return "", err
}

// UpdateFeature updates an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) UpdateFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}

// DisableFeature disables an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) DisableFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}
