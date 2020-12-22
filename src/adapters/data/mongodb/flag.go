package mongodb

import (
	"errors"

	"github.com/FeatureOn/api/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// FlagRepository represent a structure that will communicate to MongoDB to accomplish Flag related transactions
type FlagRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newFlagRepository(client *mongo.Client, databaseName string) FlagRepository {
	return FlagRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// GetFlags gets values of all active flags for a given environment
func (fr FlagRepository) GetFlags(environmentID string) ([]domain.Flag, error) {
	return nil, errors.New("Not implemented")
}

// UpdateFlag sets new value to a spesific flag
func (fr FlagRepository) UpdateFlag(productID string, environmentID string, featureID string, value bool) error {
	return errors.New("Not implemented")

}
