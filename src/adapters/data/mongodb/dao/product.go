package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductDAO represents the struct of Product type to be stored in mongoDB
type ProductDAO struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Environments []EnvironmentDAO   `bson:"environments"`
	Features     []FeatureDAO       `bson:"features"`
}

// FeatureDAO is a basic flag (as for now) holding a key within a project and its default state
type FeatureDAO struct {
	Name         string
	Key          string
	Description  string
	DefaultState bool
	Active       bool
}

// EnvironmentDAO is a struct that will hold the collection of flags for each of product's deployment
type EnvironmentDAO struct {
	ID   string
	Name string
}
