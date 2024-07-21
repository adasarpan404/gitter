package helper

import (
	"github.com/adasarpan404/gitter/database"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

var Validate = validator.New()

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
