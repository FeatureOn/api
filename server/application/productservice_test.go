package application

import (
	"fmt"
	"testing"

	"github.com/FeatureOn/api/server/domain"
	"github.com/stretchr/testify/assert"
)

type mockProductRepository struct {
}

// FlagRepository is the interface to interact with Flag domain object
type mockFlagRepository struct {
}

var (
	// GetDoFunc will be used to get different Do functions for testing purposes
	getProductByNameFunc  func(productName string) (string, error)
	addProductFunc        func(productName string) (string, error)
	updateProductFunc     func(product string, productName string) error
	getProductsFunc       func() ([]domain.Product, error)
	getProductFunc        func(id string) (domain.Product, error)
	addEnvironmentFunc    func(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error)
	updateEnvironmentFunc func(product domain.Product, environmentID string, environmentName string) error
	addFeatureFunc        func(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error
	updateFeatureFunc     func(product domain.Product, feature domain.Feature) error
	getFlagsFunc          func(environmentID string) (domain.EnvironmentFlag, error)
	updateFlagFunc        func(environmentID string, featureKey string, value bool) error
)

func (mpr mockProductRepository) GetProductByName(productName string) (string, error) {
	return getProductByNameFunc(productName)
}

func (mpr mockProductRepository) AddProduct(productName string) (string, error) {
	return addProductFunc(productName)
}
func (mpr mockProductRepository) UpdateProduct(product string, productName string) error {
	return updateProductFunc(product, productName)
}
func (mpr mockProductRepository) GetProducts() ([]domain.Product, error) {
	return getProductsFunc()
}
func (mpr mockProductRepository) GetProduct(id string) (domain.Product, error) {
	return getProductFunc(id)
}
func (mpr mockProductRepository) AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
	return addEnvironmentFunc(product, environmentName, environmentFlag)
}
func (mpr mockProductRepository) UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error {
	return updateEnvironmentFunc(product, environmentID, environmentName)
}
func (mpr mockProductRepository) AddFeature(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error {
	return addFeatureFunc(product, feature, envFlags)
}
func (mpr mockProductRepository) UpdateFeature(product domain.Product, feature domain.Feature) error {
	return updateFeatureFunc(product, feature)
}

func (mfr mockFlagRepository) GetFlags(environmentID string) (domain.EnvironmentFlag, error) {
	return getFlagsFunc(environmentID)
}
func (mfr mockFlagRepository) UpdateFlag(environmentID string, featureKey string, value bool) error {
	return updateFlagFunc(environmentID, featureKey, value)
}

func TestNewProductServiceSuccess(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	assert.NotNil(t, ps)
}

func TestGetProductByNameSuccess(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	getProductByNameFunc = func(productName string) (string, error) {
		return fmt.Sprintf("the success message: %s", productName), nil
	}
	result, err := ps.productRepository.GetProductByName("testing")
	assert.Nil(t, err)
	assert.Equal(t, result, "the success message: testing")
}
