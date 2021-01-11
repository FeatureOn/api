package application

import (
	"errors"

	"github.com/FeatureOn/api/domain"
	"github.com/rs/zerolog/log"
)

// AddFeature first checks the uniqueness of Feature's Name and Key because the system should not allow Name and Key used twice
// then adds the feature on the product to the repository injected into ProductService
func (ps ProductService) AddFeature(productID string, feature domain.Feature) error {
	// Get the product data from the data store
	product, err := ps.productRepository.GetProduct(productID)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting the product with id: %s", productID)
		return err
	}
	// Check the feature key and name is not in use
	for _, feat := range product.Features {
		if feat.Key == feature.Key {
			log.Info().Msgf("Cannot add new feature, key is not unique: %s", feature.Key)
			return errors.New("The feature key is not available")
		}
		if feat.Name == feature.Name {
			log.Info().Msgf("Cannot add new feature, name is not unique: %s", feature.Name)
			return errors.New("The feature name is not available")
		}
	}

	// Create flags for each environment
	envflags := make([]domain.EnvironmentFlag, 0)
	for _, environment := range product.Environments {
		envflag := domain.EnvironmentFlag{
			EnvironmentID: environment.ID,
		}
		flag := domain.Flag{
			FeatureKey: feature.Key,
			Value:      feature.DefaultState,
		}
		envflag.Flags = append(envflag.Flags, flag)
		envflags = append(envflags, envflag)
	}

	// Add the feature to the product
	err = ps.productRepository.AddFeature(product, feature, envflags)
	if err != nil {
		log.Error().Err(err).Msg("Error adding a new feature")
		return errors.New("Error adding a new feature")
	}
	return nil
}

// UpdateFeature first checks the uniqueness of Feature's Name and Key because the system should not allow Name and Key used twice
// then updates the feature on the product to the repository injected into ProductService
func (ps ProductService) UpdateFeature(productID string, feature domain.Feature) error {
	// Get the product data from the data store
	product, err := ps.productRepository.GetProduct(productID)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting the product with id: %s", productID)
		return err
	}
	// Check the feature name is not in use by another feature and the key exists
	found := false
	for _, feat := range product.Features {
		if feat.Name == feature.Name && feat.Key != feature.Key {
			log.Info().Msgf("Cannot update the feature, name is not unique: %s", feature.Name)
			return errors.New("The feature name is not available")
		}
		if feat.Key == feature.Key {
			found = true
		}
	}
	if !found {
		log.Info().Msgf("Cannot find the feature, productID: %s, featureKey: %s", productID, feature.Key)
		return errors.New("The feature key could not be found")
	}
	return ps.productRepository.UpdateFeature(product, feature)
}

// DisableFeature disables the specified feature on all environments of the product
func (ps ProductService) DisableFeature(product domain.Product, feat domain.Feature) error {
	return ps.productRepository.ToggleFeatureState(product, feat.Key, feat.DefaultState)
}
