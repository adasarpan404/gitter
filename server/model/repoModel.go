package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repo struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name *string            `bson:"name" validate:"required,min=2,max=100"`
}
