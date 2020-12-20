package application

import (
	"errors"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
)

// AddEnvironment first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it adds a new environment on the product to the repository injected into ProductService
func (ps ProductService) AddEnvironment(productID string, environmentName string) (string, error) {
	// Get the product data from the data store
	product, err := ps.productRepository.GetProduct(productID)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting the product with id: %s", productID)
		return "", err
	}
	// Check the environment name is not in use
	for _, env := range product.Environments {
		if env.Name == environmentName {
			log.Info().Msgf("Cannot add new environment, name is not unique: %s", environmentName)
			return "", errors.New("The environment name is not available")
		}
	}
	// Iterate through all the features to create flags for the new environment
	envflag := domain.EnvironmentFlag{}
	for _, feat := range product.Features {
		flag := domain.Flag{
			FeatureKey: feat.Key,
			Value:      feat.DefaultState,
		}
		envflag.Flags = append(envflag.Flags, flag)
	}
	environmentID, err := ps.productRepository.AddEnvironment(product, environmentName, envflag)
	if err != nil {
		log.Error().Err(err).Msg("Error adding new environment")
		return "", errors.New("Error adding new environment")
	}

	//ps.flagRepository.AddFlag(environmentID, feat.Key, feat.DefaultState)
	return environmentID, nil
}

// UpdateEnvironment first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it updates the existing environment on the product to the repository injected into ProductService
func (ps ProductService) UpdateEnvironment(productID string, environmentID string, environmentName string) error {
	// Get the product data from the data store
	product, err := ps.productRepository.GetProduct(productID)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting the product with id: %s", productID)
		return err
	}
	// Check the environment name is not in use
	for _, env := range product.Environments {
		if env.Name == environmentName {
			log.Info().Msgf("Cannot add new environment, name is not unique: %s", environmentName)
			return errors.New("The environment name is not available")
		}
	}
	return ps.productRepository.Updatenvironment(product, environmentID, environmentName)
}
