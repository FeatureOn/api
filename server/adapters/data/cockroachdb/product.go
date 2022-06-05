package cockroachdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/FeatureOn/api/server/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type ProductRepository struct {
	cp     *pgxpool.Pool
	dbName string
}

func newProductRepository(pool *pgxpool.Pool, databaseName string) ProductRepository {
	return ProductRepository{
		cp:     pool,
		dbName: databaseName,
	}
}

func (pr ProductRepository) GetProductByName(productName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := ""
	if err := pr.cp.QueryRow(ctx, fmt.Sprintf("select ID from %s.products where Name=$1", pr.dbName), productName).Scan(&id); err != nil && err.Error() != "no rows in result set" {
		return "", err
	}
	return id, nil
}

func (pr ProductRepository) AddProduct(productName string) (string, error) {
	id := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := pr.cp.Exec(ctx, fmt.Sprintf("insert into %s.products (ID, Name) values ($1, $2)", pr.dbName), id, productName)
	if err != nil {
		return "", err
	}
	if result.RowsAffected() != 1 {
		return "", errors.New("Error updating the product")
	}
	return id.String(), nil
}

func (pr ProductRepository) UpdateProduct(productID string, productName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := pr.cp.Exec(ctx, fmt.Sprintf("update %s.products set Name=$1 where ID=$2", pr.dbName), productName, productID)
	if err != nil {
		return err
	}
	if result.RowsAffected() != 1 {
		return errors.New("Error updating the product")
	}
	return nil
}

func (pr ProductRepository) GetProducts() ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := pr.cp.Query(ctx, fmt.Sprintf("select ID, Name from %s.products", pr.dbName))
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err = rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr ProductRepository) GetProduct(id string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product := domain.Product{}
	err := pr.cp.QueryRow(ctx, fmt.Sprintf("select ID, Name from %s.products where ID=$1", pr.dbName), id).Scan(&product.ID, &product.Name)
	if err != nil {
		return domain.Product{}, err
	}
	frows, err := pr.cp.Query(ctx, fmt.Sprintf("select Key, Name, Description, DefaultState, Active from %s.features where ProductID=$1", pr.dbName), id)
	defer frows.Close()
	if err != nil {
		return domain.Product{}, err
	}
	for frows.Next() {
		var feature domain.Feature
		if err := frows.Scan(&feature.Key, &feature.Name, &feature.Description, &feature.DefaultState, &feature.Active); err != nil {
			return domain.Product{}, err
		}
		product.Features = append(product.Features, feature)
	}
	erows, err := pr.cp.Query(ctx, fmt.Sprintf("select ID, Name from %s.environments where ProductID=$1", pr.dbName), id)
	defer erows.Close()
	if err != nil {
		return domain.Product{}, err
	}
	for erows.Next() {
		var environment domain.Environment
		if err := erows.Scan(&environment.ID, &environment.Name); err != nil {
			return domain.Product{}, err
		}
		product.Environments = append(product.Environments, environment)
	}
	return product, nil
}

func (pr ProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (pr ProductRepository) UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error {
	//TODO implement me
	panic("implement me")
}

func (pr ProductRepository) AddFeature(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error {
	//TODO implement me
	panic("implement me")
}

func (pr ProductRepository) UpdateFeature(product domain.Product, feature domain.Feature) error {
	//TODO implement me
	panic("implement me")
}
