package domain

// Flag represents the structure for each Feature value for each environment of a product
type Flag struct {
	EnvironmentID string
	FeatureKey    string
	Value         bool
}
