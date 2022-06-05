package cockroachdb

import (
	"github.com/FeatureOn/api/server/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FlagRepository struct {
	cp     *pgxpool.Pool
	dbName string
}

func newFlagRepository(pool *pgxpool.Pool, databaseName string) FlagRepository {
	return FlagRepository{
		cp:     pool,
		dbName: databaseName,
	}
}

func (f FlagRepository) GetFlags(environmentID string) (domain.EnvironmentFlag, error) {
	//TODO implement me
	panic("implement me")
}

func (f FlagRepository) UpdateFlag(environmentID string, featureKey string, value bool) error {
	//TODO implement me
	panic("implement me")
}
