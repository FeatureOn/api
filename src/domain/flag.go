package domain

// EnvironmentFlag respresents the structure that matches flags to the environments
type EnvironmentFlag struct {
	EnvironmentID string
	Flags         []Flag
}

// NewEnvironmentFlag respresents the structure that matches flags to the environments
type NewEnvironmentFlag struct {
	EnvironmentID string
}

// Flag represents the structure for each Feature values of environments
type Flag struct {
	FeatureKey string
	Value      bool
}
