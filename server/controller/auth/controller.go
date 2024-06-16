package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adasarpan404/gitter/helper"
	"github.com/adasarpan404/gitter/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func verifyPassword(providedPassword string, userPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		check = false
		msg = "password is incorrect"
	}
	return check, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		validationErr := helper.Validate.Struct(user)

		if validationErr != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
		}

		count, err := helper.UserCollection.CountDocuments(
			ctx,
			bson.M{
				"email": user.Email,
			})

		defer cancel()

		if err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}
		if count > 0 {
			helper.ErrorResponse(c, http.StatusInternalServerError, "this email already exists")
			return
		}

		password := HashPassword(*user.Password)

		user.Password = &password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		token, err := helper.GenerateToken(
			*user.Email,
			*user.FirstName,
			*user.LastName,
			user.ID.Hex(),
		)
		if err != nil {
			helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		userObj, err := helper.UserCollection.InsertOne(ctx, user)
		if err != nil {
			msg := "User item was not created"
			helper.ErrorResponse(c, http.StatusInternalServerError, msg)
			return
		}
		defer cancel()
		c.JSON(
			http.StatusCreated,
			gin.H{
				"status": true,
				"user":   userObj,
				"token":  token,
			})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		var user, foundUser model.User

		if err := c.BindJSON(&user); err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		err := helper.UserCollection.FindOne(
			ctx,
			bson.M{"email": user.Email},
		).Decode(&foundUser)

		defer cancel()
		if err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}
		passwordIsValid, msg := verifyPassword(*foundUser.Password, *user.Password)
		defer cancel()

		if !passwordIsValid {
			helper.ErrorResponse(c, http.StatusBadRequest, msg)
			return
		}

		if foundUser.Email == nil {
			helper.ErrorResponse(c, http.StatusBadRequest, "user not found")
			return
		}
		token, err := helper.GenerateToken(
			*foundUser.Email,
			*foundUser.FirstName,
			*foundUser.LastName,
			foundUser.ID.Hex(),
		)

		if err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": true,
				"user":   foundUser,
				"token":  token,
			})

	}
}
