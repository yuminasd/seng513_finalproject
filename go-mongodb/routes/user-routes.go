package routes

import (
	"go-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser())
	router.GET("/users/:userId", controllers.GetUser())
	router.GET("/users", controllers.GetAllUsers())

	router.POST("/groups", controllers.CreateGroup())
	router.POST("/groups/:groupId/genres", controllers.AddGenresToGroup())
	router.DELETE("/groups/:groupId/genres", controllers.DeleteGenresFromGroup())
	router.POST("/groups/:groupId/members", controllers.AddMembersToGroup())
	router.DELETE("/groups/:groupId/members/:userId", controllers.RemoveUserFromGroup())
	router.POST("/groups/:groupId/likedmovies", controllers.AddAllUsersLikedMoviesToGroup())
	router.GET("/groups/:groupId/likedmovies", controllers.GetLikedMoviesFromGroup())
	router.DELETE("/groups/:groupId/likedmovies/:movieId", controllers.DeleteLikedMovieFromGroup())
	router.POST("/groups/:groupId/likedmovies/:movieId", controllers.AddMovieToGroupLikedMovies())
	router.POST("/groups/:groupId/users/:userId/likedmovies", controllers.AddUserLikedMoviesToGroup())
}
