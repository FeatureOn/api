package application

import (
	"errors"
	"github.com/FeatureOn/api/server/domain"
	"github.com/FeatureOn/api/server/domain/mocker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddFeature(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error while checking the product
	getProductFunc = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.New("some kind of database provider error while getting the product")
	}
	mockFeature := mocker.GenerateMockFeature()
	err := ps.AddFeature("abc-def", mockFeature)
	assert.EqualError(t, err, "some kind of database provider error while getting the product")
	// Check what happens when we try to add a new feature with an existing key
	mockProduct := mocker.GenerateMockProduct(true, true)
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	mockFeature.Key = mockProduct.Features[0].Key
	err = ps.AddFeature(mockProduct.ID, mockFeature)
	assert.EqualError(t, err, "the feature key is not available")
	// Check what happens when we try to add a new feature with an existing name
	mockFeature = mocker.GenerateMockFeature()
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	mockFeature.Name = mockProduct.Features[0].Name
	err = ps.AddFeature(mockProduct.ID, mockFeature)
	assert.EqualError(t, err, "the feature name is not available")
	// Check when everything is OK but an error returned when adding the feature
	mockFeature = mocker.GenerateMockFeature()
	addFeatureFunc = func(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error {
		assert.Equal(t, len(product.Environments), len(envFlags))
		return errors.New("feature adding error")
	}
	err = ps.AddFeature(mockProduct.ID, mockFeature)
	assert.EqualError(t, err, "error adding a new feature")
	// Check when everything is OK and feature is added
	mockFeature = mocker.GenerateMockFeature()
	addFeatureFunc = func(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error {
		assert.Equal(t, len(product.Environments), len(envFlags))
		return nil
	}
	err = ps.AddFeature(mockProduct.ID, mockFeature)
	assert.Nil(t, err)
}

func TestUpdateFeature(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error while checking the product
	getProductFunc = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.New("some kind of database provider error while getting the product")
	}
	mockFeature := mocker.GenerateMockFeature()
	err := ps.UpdateFeature("abc-def", mockFeature)
	assert.EqualError(t, err, "some kind of database provider error while getting the product")
	// Check what happens when we try to update an existing feature with an existing name
	mockProduct := mocker.GenerateMockProduct(true, true)
	// Let's ensure that out mock product has more than one feature
	for len(mockProduct.Features) < 2 {
		mockProduct = mocker.GenerateMockProduct(true, true)
	}
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	mockFeature.Key = mockProduct.Features[0].Key
	mockFeature.Name = mockProduct.Features[1].Name
	err = ps.UpdateFeature(mockProduct.ID, mockFeature)
	assert.EqualError(t, err, "the feature name is not available")
	// Check what happens when we try to update a feature with non-existing key
	mockFeature = mocker.GenerateMockFeature()
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	err = ps.UpdateFeature(mockProduct.ID, mockFeature)
	assert.EqualError(t, err, "the feature key could not be found")
	// Check when everything is OK but an error returned when updating the feature
	mockFeature = mocker.GenerateMockFeature()
	updateFeatureFunc = func(product domain.Product, feature domain.Feature) error {
		return errors.New("feature updating error")
	}
	mockFeature.Key = mockProduct.Features[0].Key
	err = ps.UpdateFeature(mockProduct.ID, mockFeature)
	assert.EqualError(t, err, "feature updating error")
	// Check when everything is OK and feature is updated
	updateFeatureFunc = func(product domain.Product, feature domain.Feature) error {
		return nil
	}
	err = ps.UpdateFeature(mockProduct.ID, mockFeature)
	assert.Nil(t, err)
}
