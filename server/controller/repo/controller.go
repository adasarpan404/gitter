package repo

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
	}
}

func Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
