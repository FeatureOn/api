package application

import (
	"errors"
	"github.com/FeatureOn/api/server/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateFlagValue(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error while updating the flag
	updateFlagFunc = func(environmentID string, featureKey string, value bool) error {
		return errors.New("some kind of database provider error while updating the flag value")
	}
	err := ps.UpdateFlagValue("envID", "featureKey", true)
	assert.EqualError(t, err, "some kind of database provider error while updating the flag value")
	// Check what happens when everything is OK
	updateFlagFunc = func(environmentID string, featureKey string, value bool) error {
		return nil
	}
	err = ps.UpdateFlagValue("envID", "featureKey", true)
	assert.Nil(t, err)
}

func TestGetFlags(t *testing.T) {
	ps := NewProductService(mockProductRepository{}, mockFlagRepository{})
	// Check what happens when underlying database provider returns an error while getting the flags
	getFlagsFunc = func(environmentID string) (domain.EnvironmentFlag, error) {
		return domain.EnvironmentFlag{}, errors.New("some kind of database provider error while getting the flag value")
	}
	result, err := ps.GetFlags("envID")
	assert.Equal(t, domain.EnvironmentFlag{}, result)
	assert.EqualError(t, err, "some kind of database provider error while getting the flag value")
}
