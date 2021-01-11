package application

import "github.com/FeatureOn/api/domain"

// UpdateFeatureValue updates a Feature instance's value on its corresponding product and environment
func (ps ProductService) UpdateFeatureValue(productID string, environmentID string, featureID string, value bool) error {
	return ps.flagRepository.UpdateFlag(productID, environmentID, featureID, value)
}

// GetFlags gets all feature values for a given product and environment
func (ps ProductService) GetFlags(environmentID string) (domain.EnvironmentFlag, error) {
	return ps.flagRepository.GetFlags(environmentID)
}
