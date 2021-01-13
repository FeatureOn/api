package application

import "github.com/FeatureOn/api/server/domain"

// UpdateFlagValue updates a flag's value on its corresponding product and environment
func (ps ProductService) UpdateFlagValue(environmentID string, featureKey string, value bool) error {
	return ps.flagRepository.UpdateFlag(environmentID, featureKey, value)
}

// GetFlags gets all feature values for a given product and environment
func (ps ProductService) GetFlags(environmentID string) (domain.EnvironmentFlag, error) {
	return ps.flagRepository.GetFlags(environmentID)
}
