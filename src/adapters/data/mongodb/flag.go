package mongodb

import (
	"errors"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
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

func (fr FlagRepository) AddFlag(environmentID string, FeatureID string, value bool) error {
	return errors.New("Not implemented")
}

func (fr FlagRepository) GetFlags(environmentID string) ([]domain.Flag, error) {
	return nil, errors.New("Not implemented")
}
