package cockroachdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/FeatureOn/api/server/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
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

// GetFlags gets values of all active flags for a given environment
func (fr FlagRepository) GetFlags(environmentID string) (domain.EnvironmentFlag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := fr.cp.Query(ctx, fmt.Sprintf("select FeatureKey, Value from %s.flags where EnvironmentID=$1", fr.dbName), environmentID)
	defer rows.Close()
	envFlag := domain.EnvironmentFlag{EnvironmentID: environmentID}
	if err != nil {
		return domain.EnvironmentFlag{}, err
	}
	for rows.Next() {
		var flag domain.Flag
		if err = rows.Scan(&flag.FeatureKey, &flag.Value); err != nil {
			return domain.EnvironmentFlag{}, err
		}
		envFlag.Flags = append(envFlag.Flags, flag)
	}
	return envFlag, nil
}

// UpdateFlag sets new value to a specific flag
func (fr FlagRepository) UpdateFlag(environmentID string, featureKey string, value bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := fr.cp.Exec(ctx, fmt.Sprintf("update %s.flags set Value=$1 where FeatureKey=$2 and EnvironmentID=$3", fr.dbName), value, featureKey, environmentID)
	if err != nil {
		return err
	}
	if result.RowsAffected() != 1 {
		return errors.New("error updating the flag")
	}
	return nil
}
