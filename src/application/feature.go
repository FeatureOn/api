package application

import (
	"errors"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
)

// AddFeature first checks the uniqueness of Feature's Name and Key because the system should not allow Name and Key used twice
// then adds the feature on the product to the repository injected into ProductService
func (ps ProductService) AddFeature(productID string, feat domain.Feature) (string, error) {
	// Check the Feature's name and key are unique
	existingID, err := ps.productRepository.GetFeatureByName(productID, feat.Name)
	if err != nil {
		log.Error().Err(err).Msg("Error checking feature name uniqueness")
		return "", err
	}
	if existingID != "" {
		return "", errors.New("The feature name is not available")
	}
	existingID, err = ps.productRepository.GetFeatureByKey(productID, feat.Key)
	if err != nil {
		log.Error().Err(err).Msg("Error checking feature key uniqueness")
		return "", err
	}
	if existingID != "" {
		return "", errors.New("The feature key is not available")
	}
	// Get product's all environments
	envs, err := ps.productRepository.GetEnvironments(productID)
	if err != nil {
		log.Error().Err(err).Msg("Feature added but could not get environments")
		return "", errors.New("Error adding a new feature")
		/// ToDo: decide whether rollback or whatever
	}
	// Create flags for each environment
	envflags := make([]domain.EnvironmentFlag, 0)
	for _, environment := range envs {
		envflag := domain.EnvironmentFlag{
			EnvironmentID: environment.ID,
		}
		flag := domain.Flag{
			FeatureKey: feat.Key,
			Value:      feat.DefaultState,
		}
		envflag.Flags = append(envflag.Flags, flag)
		envflags = append(envflags, envflag)
	}

	// Add the feature to the product
	featureID, err := ps.productRepository.AddFeature(productID, feat, envflags)
	if err != nil {
		log.Error().Err(err).Msg("Error adding a new feature")
		return "", errors.New("Error adding a new feature")
	}
	return featureID, nil
}

// UpdateFeature first checks the uniqueness of Feature's Name and Key because the system should not allow Name and Key used twice
// then updates the feature on the product to the repository injected into ProductService
func (ps ProductService) UpdateFeature(productID string, feat domain.Feature) error {
	existingKey, err := ps.productRepository.GetFeatureByName(productID, feat.Name)
	if err != nil {
		log.Error().Err(err).Msg("Error checking feature name uniqueness")
		return err
	}
	if existingKey != feat.Key {
		return errors.New("The feature name is not available")
	}
	existingKey, err = ps.productRepository.GetFeatureByKey(productID, feat.Key)
	if err != nil {
		log.Error().Err(err).Msg("Error getting the feature")
		return err
	}
	if existingKey == "" {
		return errors.New("Feature key cannot be changed and there's no feature with the key provided")
	}
	return ps.productRepository.UpdateFeature(productID, feat)
}

// DisableFeature disables the specified feature on all environments of the product
func (ps ProductService) DisableFeature(productID string, feat domain.Feature) error {
	return ps.productRepository.DisableFeature(productID, feat)
}
