package application

import "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"

// UpdateFeatureValue updates a Feature instance's value on its corresponding product and environment
func (ps ProductService) UpdateFeatureValue(productID string, environmentID string, featureID string, value bool) error {
	return ps.flagRepository.UpdateFlag(productID, environmentID, featureID, value)
}

// GetValues gets all feature values for a given product and environment
func (ps ProductService) GetValues(environmentID string) ([]domain.Flag, error) {
	return ps.flagRepository.GetFlags(environmentID)
}
