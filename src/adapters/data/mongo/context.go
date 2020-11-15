package mongo

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDBContext returns a new DBContext handler with the given logger
func NewDBContext(v *Validation) *DBContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	connectionString := os.Getenv("ConnectionString")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
		log.Info().Msg("ConnectionString from Env not found, falling back to local DB")
	} else {
		log.Info().Msgf("ConnectionString from Env is used: '%s'", connectionString)
	}
	databaseName := os.Getenv("DatabaseName")
	if databaseName == "" {
		databaseName = "goboiler"
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
	return &DBContext{*client, databaseName, APIContext{v}}
}
