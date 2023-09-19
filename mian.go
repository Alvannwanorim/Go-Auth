package main

import (
	"github.com/alvannwanorim/go-auth/initializers"
	"github.com/alvannwanorim/go-auth/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}
func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	routes.AuthRoute(router)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	router.Run()
}
