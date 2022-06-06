package application

import (
	"errors"
	"github.com/FeatureOn/api/server/domain"
	"github.com/FeatureOn/api/server/domain/mocker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckProductNameSuccess(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	getProductByNameFunc = func(productName string) (string, error) {
		return "", nil
	}
	err := ps.checkProductName("testing")
	assert.Nil(t, err)
}

func TestGetProductByNameError(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when the underlying database provider returns an error
	getProductByNameFunc = func(productName string) (string, error) {
		return "", errors.New("some kind of database provider error")
	}
	err := ps.checkProductName("testing")
	assert.EqualError(t, err, "some kind of database provider error")
	// Check what happens when the underlying database provider actually return an ID
	getProductByNameFunc = func(productName string) (string, error) {
		return "abc-def", nil
	}
	err = ps.checkProductName("testing")
	assert.EqualError(t, err, "the product name is not available")
}

func TestAddProduct(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when checking the name returns error
	getProductByNameFunc = func(productName string) (string, error) {
		return "", errors.New("some kind of database provider error while checking the name")
	}
	result, err := ps.AddProduct("testing")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "some kind of database provider error while checking the name")
	// Check what happens when checking the name succeeds but adding the product returns an error
	getProductByNameFunc = func(productName string) (string, error) {
		return "", nil
	}
	addProductFunc = func(productName string) (string, error) {
		return "", errors.New("some kind of database provider error while adding the product")
	}
	result, err = ps.AddProduct("testing")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "some kind of database provider error while adding the product")
	// Check what happens when checking the name and adding the product succeeds
	getProductByNameFunc = func(productName string) (string, error) {
		return "", nil
	}
	addProductFunc = func(productName string) (string, error) {
		return "abc-def", nil
	}
	result, err = ps.AddProduct("testing")
	assert.Equal(t, "abc-def", result)
	assert.Nil(t, err)
}

func TestUpdateProduct(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when checking the name returns error
	getProductByNameFunc = func(productName string) (string, error) {
		return "", errors.New("some kind of database provider error while checking the name")
	}
	err := ps.UpdateProduct("abc-def", "testing")
	assert.EqualError(t, err, "some kind of database provider error while checking the name")
	// Check what happens when checking the name succeeds but updating the product returns an error
	getProductByNameFunc = func(productName string) (string, error) {
		return "", nil
	}
	updateProductFunc = func(productID string, productName string) error {
		return errors.New("some kind of database provider error while updating the product")
	}
	err = ps.UpdateProduct("abc-def", "testing")
	assert.EqualError(t, err, "some kind of database provider error while updating the product")
	// Check what happens when checking the name and updating the product succeeds
	getProductByNameFunc = func(productName string) (string, error) {
		return "", nil
	}
	updateProductFunc = func(productID string, productName string) error {
		return nil
	}
	err = ps.UpdateProduct("abc-def", "testing")
	assert.Nil(t, err)
}

func TestGetProduct(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error
	getProductFunc = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.New("some kind of database provider error while getting the product")
	}
	result, err := ps.GetProduct("abc-def")
	assert.Equal(t, domain.Product{}, result)
	assert.EqualError(t, err, "some kind of database provider error while getting the product")
	// Check what happens when underlying database provider returns a product
	mockproduct := mocker.GenerateMockProduct()
	getProductFunc = func(id string) (domain.Product, error) {
		return mockproduct, nil
	}
	result, err = ps.GetProduct("abc-def")
	assert.Equal(t, mockproduct, result)
	assert.Nil(t, err)
}
