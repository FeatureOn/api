package cockroachdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/FeatureOn/api/server/domain"
	"github.com/google/uuid"
	"time"
)

// AddEnvironment adds a new environment together with all its flags with default values and returns its ID,
// returns empty string and error otherwise
func (pr ProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	id := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tran, err := pr.cp.Begin(ctx)
	if err != nil {
		return "", err
	}
	result, err := pr.cp.Exec(ctx, fmt.Sprintf("insert into %s.environments (ID, Name, ProductID) values ($1, $2, $3)", pr.dbName), id, environmentName, product.ID)
	if err != nil {
		tran.Rollback(ctx)
		return "", err
	}
	if result.RowsAffected() != 1 {
		tran.Rollback(ctx)
		return "", errors.New("error adding the environment")
	}
	// We have to add flags for this environment for each feature
	for _, flag := range environmentFlag.Flags {
		result, err = pr.cp.Exec(ctx, fmt.Sprintf("insert into %s.flags (FeatureKey, EnvironmentID, Value) values ($1, $2, $3)", pr.dbName), flag.FeatureKey, id, flag.Value)
		if err != nil {
			tran.Rollback(ctx)
			return "", err
		}
		if result.RowsAffected() != 1 {
			tran.Rollback(ctx)
			return "", errors.New("error adding the environment")
		}
	}
	tran.Commit(ctx)
	return id.String(), nil
}

// UpdateEnvironment updates an existing environment on the database
func (pr ProductRepository) UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := pr.cp.Exec(ctx, fmt.Sprintf("update %s.environments set Name=$1 where ID=$2", pr.dbName), environmentName, environmentID)
	if err != nil {
		return err
	}
	if result.RowsAffected() != 1 {
		return errors.New("error updating the environment")
	}
	return nil
}
