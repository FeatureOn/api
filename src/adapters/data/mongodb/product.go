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
)

// ProductRepository represent a structpre that will communicate to MongoDB to accomplish product related transactions
type ProductRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newProductRepository(client *mongo.Client, databaseName string) ProductRepository {
	return ProductRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// GetProduct retprns one Product with the given ID if it exists in the array, returns not found error otherwise
func (pr ProductRepository) GetProduct(ID string) (domain.Product, error) {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error().Err(err).Msgf("Error parsing ProductID: %s", ID)
		return domain.Product{}, err
	}
	var ProductDAO dao.ProductDAO
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&ProductDAO)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting Product with ProductID: %s", ID)
		return domain.Product{}, err
	}
	return mappers.MapProductDAO2Product(ProductDAO), nil
}

// GetProducts returns an array of all products defined in the database
func (pr ProductRepository) GetProducts() ([]domain.Product, error) {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var productDAO dao.ProductDAO
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Error().Err(err).Msgf("Error getting Products")
		return nil, err
	}
	products := make([]domain.Product, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err := cur.Decode(&productDAO)
		if err != nil {
			return nil, err
		}
		product := mappers.MapProductDAO2Product(productDAO)
		products = append(products, product)
	}
	return products, nil
}

// GetProductByName returnd the ID of the product if the name matches a product in the database, returns empty string and error otherwise
func (pr ProductRepository) GetProductByName(productName string) (string, error) {
	return "", errors.New("Not implemented")
}

// AddProduct adds a new product to the database and returns its ID, returns empty string and error otherwise
func (pr ProductRepository) AddProduct(productName string) (string, error) {
	return "", errors.New("Not implemented")

}

// UpdateProduct updates a product on the database, returns error otherwise
func (pr ProductRepository) UpdateProduct(productID string, productName string) error {
	return errors.New("Not implemented")

}

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
		} else {
			log.Error().Err(err).Msgf("The productID %s did not match any products in the database", product.ID)
			return "", errors.New("Product not found")
		}
	} else {
		log.Error().Err(err).Msgf("Error adding feature with key %s", newFeat.Key)
		return "", err
	}
}

// UpdateFeature updates an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) UpdateFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}

// DisableFeature disables an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) DisableFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")
}
