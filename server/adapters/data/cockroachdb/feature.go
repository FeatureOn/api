package cockroachdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/FeatureOn/api/server/domain"
	"time"
)

// AddFeature adds a new feature to an existing product on the database together with flags for all environments
// od the product with default values. Returns ID if successful, empty string and error otherwise
func (pr ProductRepository) AddFeature(product domain.Product, feature domain.Feature, environmentFlags []domain.EnvironmentFlag) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	tran, err := pr.cp.Begin(ctx)
	if err != nil {
		return err
	}
	result, err := pr.cp.Exec(ctx, fmt.Sprintf("insert into %s.features (Key, Name, ProductID, Description, DefaultState, Active) values ($1, $2, $3, $4, $5, $6)", pr.dbName), feature.Key, feature.Name, product.ID, feature.Description, feature.DefaultState, feature.Active)
	if err != nil {
		tran.Rollback(ctx)
		return err
	}
	if result.RowsAffected() != 1 {
		tran.Rollback(ctx)
		return errors.New("error adding the feature")
	}
	// We have to add flags for this feature for each environment
	for _, flag := range environmentFlags {
		result, err = pr.cp.Exec(ctx, fmt.Sprintf("insert into %s.flags (FeatureKey, EnvironmentID, Value) values ($1, $2, $3)", pr.dbName), feature.Key, flag.EnvironmentID, feature.DefaultState)
		if err != nil {
			tran.Rollback(ctx)
			return err
		}
		if result.RowsAffected() != 1 {
			tran.Rollback(ctx)
			return errors.New("error adding the feature")
		}
	}
	tran.Commit(ctx)
	return nil
}

// UpdateFeature updates an existing Feature on the database. Returns error if not successful
func (pr ProductRepository) UpdateFeature(product domain.Product, feature domain.Feature) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := pr.cp.Exec(ctx, fmt.Sprintf("update %s.features set Name=$1, Description=$2, DefaultState=$3, Active=$4 where Key=$5 and ProductID=$6", pr.dbName), feature.Name, feature.Description, feature.DefaultState, feature.Active, feature.Key, product.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() != 1 {
		return errors.New("error updating the feature")
	}
	return nil
}
