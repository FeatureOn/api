package dao

// EnvironmentFlag respresents the structure that matches flags to the environments
type EnvironmentFlag struct {
	EnvironmentID string `bson:"environmentID"`
	Flags         []Flag `bson:"flags"`
}

// Flag represents the structure for each Feature values of environments
type Flag struct {
	FeatureKey string `bson:"featureKey"`
	Value      bool   `bson:"value"`
}
