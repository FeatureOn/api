package application

import (
	"errors"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
)

// AddEnvironment first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it adds a new environment on the product to the repository injected into ProductService
func (ps ProductService) AddEnvironment(productID string, environmentName string) (string, error) {
	existingID, err := ps.productRepository.GetEnvironmentByName(productID, environmentName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking environment name uniqueness")
		return "", err
	}
	if existingID != productID {
		return "", errors.New("The environment name is not available")
	}
	environmentID, err := ps.productRepository.AddEnvironment(productID, environmentName)
	if err != nil {
		log.Error().Err(err).Msg("Error adding new environment")
		return "", errors.New("Error adding new environment")
	}
	/// ToDo: Add code to add all active flags to new environment
	features, err := ps.productRepository.GetFeatures(productID)
	for _, feat := range features {
		ps.flagRepository.AddFlag(environmentID, feat.Key, feat.DefaultState)
	}
	return environmentID, nil
}

// UpdateEnvironment first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it updates the existing environment on the product to the repository injected into ProductService
func (ps ProductService) UpdateEnvironment(productID string, environmentID string, environmentName string) error {
	existingID, err := ps.productRepository.GetEnvironmentByName(productID, environmentName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking environment name uniqueness")
		return err
	}
	if existingID != environmentName {
		return errors.New("The environment name is not available")
	}
	return ps.productRepository.Updatenvironment(productID, environmentID, environmentName)
}

// GetEnvironments returns a the list of environments defined for a product from the repository injected into ProductService
func (ps ProductService) GetEnvironments(productID string) ([]domain.Environment, error) {
	return ps.productRepository.GetEnvironments(productID)
}

// GetEnvironment returns an environment defined for a product from the repository injected into ProductService
func (ps ProductService) GetEnvironment(productID string, environmentID string) (domain.Environment, error) {
	return ps.productRepository.GetEnvironment(productID, environmentID)
}
