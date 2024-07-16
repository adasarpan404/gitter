package userController

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adasarpan404/gitter/helper"
	"github.com/adasarpan404/gitter/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Get("userId")
		if !ok {
			helper.ErrorResponse(c, http.StatusBadRequest, "User ID not found in context")
			return
		}

		objectUserId, err := primitive.ObjectIDFromHex(fmt.Sprint(userId))
		if err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID format")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		var user model.User

		projection := primitive.M{"password": 0}

		err = helper.UserCollection.FindOne(ctx,
			bson.M{
				"_id": objectUserId,
			},
			options.FindOne().SetProjection(projection),
		).Decode(&user)

		if err != nil {
			helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": true,
				"user":   user,
			})
	}
}
