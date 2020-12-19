package dto

// EnvironmentFlagResponse respresents the structure that returns from the rest service
type EnvironmentFlagResponse struct {
	EnvironmentID string
	Flags         []FlagResponse
}

// FlagResponse represents the structure for each Feature values of environments
type FlagResponse struct {
	FeatureKey string
	Value      bool
}
