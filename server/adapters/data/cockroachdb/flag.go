package cockroachdb

import (
	"github.com/FeatureOn/api/server/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FlagRepository struct {
	cp *pgxpool.Pool
}

func newFlagRepository(pool *pgxpool.Pool) FlagRepository {
	return FlagRepository{
		cp: pool,
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
