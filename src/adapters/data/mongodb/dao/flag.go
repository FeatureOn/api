package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

// EnvironmentFlagDAO respresents the structure that matches flags to the environments
type EnvironmentFlagDAO struct {
	EnvironmentID primitive.ObjectID `bson:"environmentID"`
	Flags         []FlagDAO          `bson:"flags"`
}

// FlagDAO represents the structure for each Feature values of environments
type FlagDAO struct {
	FeatureKey string `bson:"featureKey"`
	Value      bool   `bson:"value"`
}
