package dto

// EnvironmentFlagResponse respresents the structure that returns from the rest service
type EnvironmentFlagResponse struct {
	EnvironmentID string         `json:"environmentID"`
	Flags         []FlagResponse `json:"flags"`
}

// FlagResponse represents the structure for each Feature values of environments
type FlagResponse struct {
	FeatureKey string `json:"featureKey"`
	Value      bool   `json:"value"`
}

// UpdateFlagRequest represents teh structure that will be used to update value of a flag
type UpdateFlagRequest struct {
	EnvironmentID string `json:"environmentID" validate:"required"`
	FeatureKey    string `json:"featureKey" validate:"required"`
	Value         bool   `json:"value"`
}
