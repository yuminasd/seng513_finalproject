package controllers

import (
	"context"
	"log"
	"go-mongodb/configs"
	"go-mongodb/models"
	"go-mongodb/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var groupCollection *mongo.Collection = configs.GetCollection(configs.DB, "groups")

func AddGenresToGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupID := c.Param("groupId")
		defer cancel()

		var group models.Group

		objID, err := primitive.ObjectIDFromHex(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": "Invalid group ID"},
			})
			return
		}

		err = groupCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var genresToAdd []string
		if err := c.BindJSON(&genresToAdd); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Check if genresToAdd already exist in the group's genres
		for _, newGenre := range genresToAdd {
			exists := false
			for _, existingGenre := range group.Genre {
				if newGenre == existingGenre {
					exists = true
					break
				}
			}
			if !exists {
				group.Genre = append(group.Genre, newGenre)
			}
		}

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Genres successfully added to the group!"},
		})
	}
}

func DeleteGenresFromGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupID := c.Param("groupId") // Change the parameter name to "id"
		defer cancel()

		var group models.Group
		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		var genresToDelete []string
		if err := c.BindJSON(&genresToDelete); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group) // Change the query to use "id"
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		for _, genre := range genresToDelete {
			for i, g := range group.Genre {
				if g == genre {
					group.Genre = append(group.Genre[:i], group.Genre[i+1:]...)
					break
				}
			}
		}

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group) // Change the query to use "id"
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Genres removed from the group"},
		})
	}
}

func AddMembersToGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupID := c.Param("groupId")
		defer cancel()

		var group models.Group

		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var memberIDs []string
		if err := c.BindJSON(&memberIDs); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		for _, memberID := range memberIDs {
			var user models.User
			userID, _ := primitive.ObjectIDFromHex(memberID)
			err := userCollection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
			group.Members = append(group.Members, user)
		}

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Members successfully added to the group!"},
		})
	}
}

func RemoveUserFromGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupID := c.Param("groupId")
		userID := c.Param("userId")
		defer cancel()

		var group models.Group
		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		userIndex := -1
		for i, user := range group.Members {
			if user.Id.Hex() == userID {
				userIndex = i
				break
			}
		}

		if userIndex == -1 {
			c.JSON(http.StatusNotFound, responses.UserResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    map[string]interface{}{"data": "User not found in the group"},
			})
			return
		}

		group.Members = append(group.Members[:userIndex], group.Members[userIndex+1:]...)

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "User successfully removed from the group"},
		})
	}
}

func AddAllUsersLikedMoviesToGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		groupID := c.Param("groupId")
		var group models.Group
		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		for _, member := range group.Members {
			var user models.User
			err := userCollection.FindOne(ctx, bson.M{"id": member.Id}).Decode(&user)
			if err != nil {
				log.Println("Error finding user for ID:", member.Id)
				continue
			}

			log.Println("User ID:", user.Id)
			log.Println("User Liked Movies:", user.LikedMovies)

			if len(user.LikedMovies) == 0 {
				log.Println("No liked movies for user ID:", user.Id)
				continue
			}

			for _, likedMovieID := range user.LikedMovies {
				log.Println("Checking movies for user ID:", user.Id)
				log.Println("Liked Movie ID:", likedMovieID)
				objLikedMovieID, _ := primitive.ObjectIDFromHex(likedMovieID)

				var movie models.Movie
				err := movieCollection.FindOne(ctx, bson.M{"id": objLikedMovieID}).Decode(&movie)
				if err != nil {
					log.Println("Error finding movie for ID:", objLikedMovieID)
					continue
				}

				log.Println("Movie Genres:", movie.Genres)

				for _, movieGenre := range movie.Genres {
					for _, groupGenre := range group.Genre {
						if movieGenre == groupGenre {
							log.Println("Movie Matched Genre:", movie.Id, movieGenre)
							found := false
							for idx, likedMovie := range group.LikedMovies {
								if likedMovie.MovieId == likedMovieID {
									group.LikedMovies[idx].LikedCount++
									found = true
									break
								}
							}
							if !found {
								group.LikedMovies = append(group.LikedMovies, models.Likedmovies{
									MovieId:    likedMovieID,
									LikedCount: 1,
								})
							}
							break 
						}
					}
				}
			}
		}

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Liked movies added to the group"},
		})
	}
}

func GetLikedMoviesFromGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		groupID := c.Param("groupId")
		var group models.Group
		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"likedMovies": group.LikedMovies},
		})
	}
}

func DeleteLikedMovieFromGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		groupID := c.Param("groupId")
		movieID := c.Param("movieId")

		var group models.Group
		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		for idx, likedMovie := range group.LikedMovies {
			if likedMovie.MovieId == movieID {
				if likedMovie.LikedCount > 1 {
					group.LikedMovies[idx].LikedCount--
				} else {
					group.LikedMovies = append(group.LikedMovies[:idx], group.LikedMovies[idx+1:]...)
				}
				break
			}
		}

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Liked movie deleted from the group"},
		})
	}
}

func AddMovieToGroupLikedMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		groupID := c.Param("groupId")
		movieID := c.Param("movieId")

		var group models.Group
		objGroupID, err := primitive.ObjectIDFromHex(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid group ID",
			})
			return
		}

		err = groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error finding group",
				"error":   err.Error(),
			})
			return
		}

		for idx, likedMovie := range group.LikedMovies {
			if likedMovie.MovieId == movieID {
				group.LikedMovies[idx].LikedCount++
				_, err := groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Error updating group",
						"error":   err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"message": "Liked count increased for the movie in the group",
				})
				return
			}
		}

		group.LikedMovies = append(group.LikedMovies, models.Likedmovies{
			MovieId:    movieID,
			LikedCount: 1,
		})
		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating group",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Movie added to the group's liked movies",
		})
	}
}

func AddUserLikedMoviesToGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		groupID := c.Param("groupId")
		userID := c.Param("userId") 
		var group models.Group
		objGroupID, _ := primitive.ObjectIDFromHex(groupID)

		err := groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var user models.User
		objUserID, _ := primitive.ObjectIDFromHex(userID)
		err = userCollection.FindOne(ctx, bson.M{"id": objUserID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		log.Println("User ID:", user.Id)
		log.Println("User Liked Movies:", user.LikedMovies)

		if len(user.LikedMovies) == 0 {
			log.Println("No liked movies for user ID:", user.Id)
			return
		}

		for _, likedMovieID := range user.LikedMovies {
			log.Println("Checking movies for user ID:", user.Id)
			log.Println("Liked Movie ID:", likedMovieID)
			objLikedMovieID, _ := primitive.ObjectIDFromHex(likedMovieID)

			var movie models.Movie
			err := movieCollection.FindOne(ctx, bson.M{"id": objLikedMovieID}).Decode(&movie)
			if err != nil {
				log.Println("Error finding movie for ID:", objLikedMovieID)
				continue
			}

			log.Println("Movie Genres:", movie.Genres)

			for _, movieGenre := range movie.Genres {
				for _, groupGenre := range group.Genre {
					if movieGenre == groupGenre {
						log.Println("Movie Matched Genre:", movie.Id, movieGenre)
						found := false
						for idx, likedMovie := range group.LikedMovies {
							if likedMovie.MovieId == likedMovieID {
								group.LikedMovies[idx].LikedCount++
								found = true
								break
							}
						}
						if !found {
							group.LikedMovies = append(group.LikedMovies, models.Likedmovies{
								MovieId:    likedMovieID,
								LikedCount: 1,
							})
						}
						break
					}
				}
			}
		}

		_, err = groupCollection.ReplaceOne(ctx, bson.M{"id": objGroupID}, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Liked movies added to the group"},
		})
	}
}


