package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductDAO represents the struct of Product type to be stored in mongoDB
type ProductDAO struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
