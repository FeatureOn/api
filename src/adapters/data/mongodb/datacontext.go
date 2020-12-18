package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	UserRepository    UserRepository
	HealthRepository  HealthRepository
	ProductRepository ProductRepository
	FlagRepository    FlagRepository
}

// NewDataContext returns a new mongoDB backed DataContext
func NewDataContext() DataContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	connectionString := os.Getenv("ConnectionString")
	if connectionString == "" {
		connectionString = "mongodb+srv://toggleruser:toggpassler@toggler.au80d.mongodb.net/<dbname>?retryWrites=true&w=majority"
		log.Info().Msg("ConnectionString from Env not found, falling back to local DB")
	} else {
		log.Info().Msgf("ConnectionString from Env is used: '%s'", connectionString)
	}
	databaseName := os.Getenv("DatabaseName")
	if databaseName == "" {
		databaseName = "toggler"
		log.Info().Msg("DatabaseName from Env not found, falling back to default")
	} else {
		log.Info().Msgf("DatabaseName from Env is used: '%s'", databaseName)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("An error occured while connecting to tha database")
	} else {

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Error().Err(err).Msg("An error occured while connecting to tha database")
		}
		log.Info().Msg("Connected to MongoDB!")
	}
	dataContext := DataContext{}
	dataContext.UserRepository = newUserRepository(client, databaseName)
	dataContext.HealthRepository = newHealthRepository(client, databaseName)
	dataContext.ProductRepository = newProductRepository(client, databaseName)
	dataContext.FlagRepository = newFlagRepository(client, databaseName)
	return dataContext
}
