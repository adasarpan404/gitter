package routes

import (
	authController "github.com/adasarpan404/gitter/controller/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/login", authController.Login())
	incomingRoutes.POST("/auth/signup", authController.Signup())
}
