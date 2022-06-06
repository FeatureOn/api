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

// GetProduct returns one Product with the given ID if it exists in the array, returns not found error otherwise
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

// GetProductByName returns the ID of the product if the name matches a product in the database, returns empty string and error otherwise
func (pr ProductRepository) GetProductByName(productName string) (string, error) {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var productDAO dao.NewProductDAO
	err := collection.FindOne(ctx, bson.M{"name": productName}, options.FindOne().SetProjection(bson.M{"_id": 1})).Decode(&productDAO)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return "", errors.New("error checking product name authenticity")
	}
	return productDAO.ID.Hex(), nil
}

// AddProduct adds a new product to the database and returns its ID, returns empty string and error otherwise
func (pr ProductRepository) AddProduct(productName string) (string, error) {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	newProdID := primitive.NewObjectID()
	newProduct := dao.NewProductDAO{
		ID:   newProdID,
		Name: productName,
	}
	_, err := collection.InsertOne(ctx, newProduct)
	if err != nil {
		log.Error().Err(err).Msgf("Error adding product with name %s", productName)
		return "", errors.New("error adding product")
	}
	return newProdID.Hex(), nil

}

// UpdateProduct updates a product on the database, returns error otherwise
func (pr ProductRepository) UpdateProduct(productID string, productName string) error {
	collection := pr.dbClient.Database(pr.dbName).Collection(viper.GetString("ProductsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse productID: %s into ObjectID", productID)
		return errors.New("ProductID format is not as expected")
	}
	idDoc := bson.D{{Key: "_id", Value: id}}
	upDoc := bson.D{{Key: "$set", Value: bson.M{"name": productName}}}
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	_, err = collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
	if err != nil {
		log.Error().Err(err).Msgf("Error updating the product with productID: %s", productID)
		return errors.New("error updating the product")
	}
	return nil
}
