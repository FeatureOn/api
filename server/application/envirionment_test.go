package application

import (
	"errors"
	"github.com/FeatureOn/api/server/domain"
	"github.com/FeatureOn/api/server/domain/mocker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddEnvironment(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error while checking the product
	getProductFunc = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.New("some kind of database provider error while getting the product")
	}
	mockEnvironment := mocker.GenerateMockEnvironment()
	result, err := ps.AddEnvironment("abc-def", mockEnvironment.Name)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "some kind of database provider error while getting the product")
	// Check what happens when we try to add a new feature with an existing name
	mockProduct := mocker.GenerateMockProduct(true, true)
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	mockEnvironment.Name = mockProduct.Environments[0].Name
	result, err = ps.AddEnvironment(mockProduct.ID, mockEnvironment.Name)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "the environment name is not available")
	// Check when everything is OK but an error returned when adding the feature
	mockEnvironment = mocker.GenerateMockEnvironment()
	addEnvironmentFunc = func(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
		assert.Equal(t, len(product.Features), len(environmentFlag.Flags))
		return "", errors.New("environment adding error")
	}
	result, err = ps.AddEnvironment(mockProduct.ID, mockEnvironment.Name)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "error adding new environment")
	// Check when everything is OK and feature is added
	mockEnvironment = mocker.GenerateMockEnvironment()
	addEnvironmentFunc = func(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error) {
		assert.Equal(t, len(product.Features), len(environmentFlag.Flags))
		return "newEnvID", nil
	}
	result, err = ps.AddEnvironment(mockProduct.ID, mockEnvironment.Name)
	assert.Equal(t, "newEnvID", result)
	assert.Nil(t, err)
}

func TestUpdateEnvironment(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error while checking the product
	getProductFunc = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.New("some kind of database provider error while getting the product")
	}
	mockEnvironment := mocker.GenerateMockEnvironment()
	err := ps.UpdateEnvironment("abc-def", mockEnvironment.ID, mockEnvironment.Name)
	assert.EqualError(t, err, "some kind of database provider error while getting the product")
	// Check what happens when we try to update an existing environment with an existing name
	mockProduct := mocker.GenerateMockProduct(true, true)
	// Let's ensure that out mock product has more than one Environment
	for len(mockProduct.Environments) < 2 {
		mockProduct = mocker.GenerateMockProduct(true, true)
	}
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	mockEnvironment.ID = mockProduct.Environments[0].ID
	mockEnvironment.Name = mockProduct.Environments[1].Name
	err = ps.UpdateEnvironment(mockProduct.ID, mockEnvironment.ID, mockEnvironment.Name)
	assert.EqualError(t, err, "the environment name is not available")
	// Check what happens when we try to update a Environment with non-existing key
	mockEnvironment = mocker.GenerateMockEnvironment()
	getProductFunc = func(id string) (domain.Product, error) {
		return mockProduct, nil
	}
	err = ps.UpdateEnvironment(mockProduct.ID, mockEnvironment.ID, mockEnvironment.Name)
	assert.EqualError(t, err, "the environment id could not be found")
	// Check when everything is OK but an error returned when updating the Environment
	mockEnvironment = mocker.GenerateMockEnvironment()
	updateEnvironmentFunc = func(product domain.Product, environmentID string, environmentName string) error {
		return errors.New("environment updating error")
	}
	mockEnvironment.ID = mockProduct.Environments[0].ID
	err = ps.UpdateEnvironment(mockProduct.ID, mockEnvironment.ID, mockEnvironment.Name)
	assert.EqualError(t, err, "environment updating error")
	// Check when everything is OK and Environment is updated
	updateEnvironmentFunc = func(product domain.Product, environmentID string, environmentName string) error {
		return nil
	}
	err = ps.UpdateEnvironment(mockProduct.ID, mockEnvironment.ID, mockEnvironment.Name)
	assert.Nil(t, err)
}
