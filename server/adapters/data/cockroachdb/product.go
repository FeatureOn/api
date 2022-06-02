package cockroachdb

import (
	"github.com/FeatureOn/api/server/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepository struct {
	cp *pgxpool.Pool
}

func newProductRepository(pool *pgxpool.Pool) ProductRepository {
	return ProductRepository{
		cp: pool,
	}
}

func (p ProductRepository) GetProductByName(productName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) AddProduct(productName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateProduct(product string, productName string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetProducts() ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetProduct(id string) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) AddFeature(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateFeature(product domain.Product, feature domain.Feature) error {
	//TODO implement me
	panic("implement me")
}
