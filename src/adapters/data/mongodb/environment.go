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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddEnvironment adds a new environment together with all its flags with default values and returns its ID,
// returns empty string and error otherwise
func (pr ProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	productDAO := mappers.MapProduct2ProductDAO(product)
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	newEnvID := primitive.NewObjectID()
	newEnv := dao.EnvironmentDAO{
		ID:   newEnvID,
		Name: environmentName,
	}
	productDAO.Environments = append(productDAO.Environments, newEnv)
	idDoc := bson.D{{"_id", productDAO.ID}}

	upDoc := bson.D{{"$set", bson.M{"environments": productDAO.Environments}}}
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	result, err := collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
	if err == nil {
		if result.MatchedCount == 1 {
			// Add the flags for the new environment
			flagCollection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("FlagsCollection"))
			environmentFlag.EnvironmentID = newEnvID.Hex()
			_, err = flagCollection.InsertOne(ctx, mappers.MapEnvironmentFlag2EnvironmentFlagDAO(environmentFlag))
			return newEnvID.Hex(), nil
		} else {
			log.Error().Err(err).Msgf("The productID %s did not match any products in the database", product.ID)
			return "", errors.New("Product not found")
		}
	} else {
		log.Error().Err(err).Msgf("Error adding environment with name %s", environmentName)
		return "", err
	}
}

// UpdateEnvironment updates an existing environment on the database
func (pr ProductRepository) UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error {
	return errors.New("Not implemented")
}
