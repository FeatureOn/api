package mongodb

import (
	"context"
	"errors"
	"time"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/data/mongodb/dao"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/data/mongodb/mappers"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// GetProduct retprns one Product with the given ID if it exists in the array, retprns not found error otherwise
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

func (pr ProductRepository) GetProductByName(productName string) (string, error) {
	return "", errors.New("Not implemented")
}
func (pr ProductRepository) AddProduct(productName string) (string, error) {
	return "", errors.New("Not implemented")

}
func (pr ProductRepository) UpdateProduct(productID string, productName string) error {
	return errors.New("Not implemented")

}

func (pr ProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	return "", errors.New("Not implemented")

}

func (pr ProductRepository) Updatenvironment(product domain.Product, environmentID string, environmentName string) error {
	return errors.New("Not implemented")

}

func (pr ProductRepository) AddFeature(product domain.Product, feat domain.Feature, envFlags []domain.EnvironmentFlag) (string, error) {
	return "", errors.New("Not implemented")

}
func (pr ProductRepository) UpdateFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")

}
func (pr ProductRepository) DisableFeature(product domain.Product, feat domain.Feature) error {
	return errors.New("Not implemented")

}
