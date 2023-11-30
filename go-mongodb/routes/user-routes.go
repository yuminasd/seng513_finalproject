package routes

import (
	"go-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users/:userId", controllers.GetUser())
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/users/liked/:userId", controllers.GetLiked())
	router.GET("/users/disliked/:userId", controllers.GetDisliked())
	router.GET("/movies", controllers.GetAllMovies())
	router.GET("/movies/:movieId", controllers.GetMovie())
	router.GET("/moviesbygenre", controllers.GetMoviesByGenre())
	router.GET("/checklogin", controllers.CheckAuth())

	router.POST("/users", controllers.CreateUser())
	router.POST("/movies", controllers.CreateMovie())

	router.DELETE("/movies/:movieId", controllers.DeleteMovie())
	router.DELETE("/users/:userId", controllers.DeleteAUser())

	router.PUT("/users/:userId", controllers.UpdateUser())

	router.PATCH("/removegroup/:userId/:groupId", controllers.RemoveGroup())
	router.PATCH("/addgroup/:userId/:groupId", controllers.AddGroup())
	router.PATCH("/addliked/:userId/:movieId", controllers.AddLiked())
	router.PATCH("/removeliked/:userId/:movieId", controllers.RemoveLiked())
	router.PATCH("/adddisliked/:userId/:movieId", controllers.AddDisliked())
	router.PATCH("/removedisliked/:userId/:movieId", controllers.RemoveDisliked())
}
