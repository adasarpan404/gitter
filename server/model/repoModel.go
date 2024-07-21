package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `bson:"name" validate:"required,min=2,max=100"`
	User      primitive.ObjectID `bson:"user" `
	CreatedAt time.Time          `bson:"createdAt"`
	UpdateAt  time.Time          `bson:"updatedAt"`
}
