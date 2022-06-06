package cockroachdb

import (
	"context"
	"github.com/FeatureOn/api/server/adapters/data"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// NewDataContext returns a new CockroachDB/Postgres backed DataContext
func NewDataContext() data.DataContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	connectionString := os.Getenv("ConnectionString")
	if connectionString == "" {
		connectionString = "postgresql://root@127.0.0.1:26257/defaultdb?sslmode=disable"
		log.Info().Msg("ConnectionString from Env not found, falling back to local DB")
	} else {
		log.Info().Msgf("ConnectionString from Env is used: '%s'", connectionString)
	}
	databaseName := os.Getenv("DatabaseName")
	if databaseName == "" {
		databaseName = "featureon"
		log.Info().Msg("DatabaseName from Env not found, falling back to default")
	} else {
		log.Info().Msgf("DatabaseName from Env is used: '%s'", databaseName)
	}
	dbPool, err := pgxpool.Connect(ctx, connectionString)
	// execute the select query and get result rows
	_, err = dbPool.Query(ctx, "select 1 from featureon.users")

	if err != nil {
		log.Error().Err(err).Msgf("An error occured while connecting to tha database")
	}
	dataContext := data.DataContext{}
	dataContext.UserRepository = newUserRepository(dbPool, databaseName)
	dataContext.HealthRepository = newHealthRepository(dbPool, databaseName)
	dataContext.ProductRepository = newProductRepository(dbPool, databaseName)
	dataContext.FlagRepository = newFlagRepository(dbPool, databaseName)
	return dataContext
}
