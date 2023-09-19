package routes

import (
	"github.com/alvannwanorim/go-auth/controllers"
	"github.com/alvannwanorim/go-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoute(route *gin.Engine) {
	route.POST("/sign-up", controllers.CreateUser)
	route.POST("/login", controllers.Login)
	route.GET("/validate", middlewares.AuthMiddleware, controllers.Validate)
}
