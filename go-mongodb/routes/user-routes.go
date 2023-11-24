package routes

import (
	"go-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser())
	router.GET("/users/:userId", controllers.GetUser())
	router.GET("/users", controllers.GetAllUsers())
	router.DELETE("users/:userId", controllers.DeleteAUser())
}
