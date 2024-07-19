package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Branch struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `bson:"name" validate:"required,min=2,max=100"`
	Repo      primitive.ObjectID `bson:"repo"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
