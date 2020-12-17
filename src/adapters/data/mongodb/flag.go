package mongodb

import (
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
