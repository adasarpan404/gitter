package main

import (
	"github.com/adasarpan404/gitter/environment"
	"github.com/adasarpan404/gitter/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Run(":" + environment.PORT)
}
