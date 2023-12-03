package controllers

import (
	"context"
	"go-mongodb/configs"
	"go-mongodb/models"
	"go-mongodb/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var groupCollection *mongo.Collection = configs.GetCollection(configs.DB, "groups")
var validateG = validator.New()

func updateGroup(groupID primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var group models.Group

	err := groupCollection.FindOne(ctx, bson.M{"id": groupID}).Decode(&group)
	if err != nil {
		return
	}

	var newUserList []models.User
	var newUser models.User
	for i := 0; i < len(group.Members); i++ {
		groupID := group.Members[i].Id
		updateUserGroups(group.Members[i].Id)
		err := userCollection.FindOne(ctx, bson.M{"id": groupID}).Decode(&newUser)
		if err != nil {
			continue
		}
		newUser.GroupID = nil
		newUserList = append(newUserList, newUser)
	}

	group.Members = newUserList
	userCollection.ReplaceOne(ctx, bson.M{"id": groupID}, group)
}

type GroupIdReturn struct {
	InsertedID primitive.ObjectID
}

func CreateGroup() gin.HandlerFunc { //Should probably check if it exists already
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var group models.Group
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateG.Struct(&group); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newGroup := models.Group{
			Id:          primitive.NewObjectID(),
			Name:        group.Name,
			Genre:       group.Genre,
			Members:     group.Members,
			LikedMovies: group.LikedMovies,
		}

		_, err := groupCollection.InsertOne(ctx, newGroup)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var results GroupIdReturn
		results.InsertedID = newGroup.Id
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": results}})
	}
}

// func GetGroupInfo(id) / get all name, users, genre
func GetGroupInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupId := c.Param("groupId")
		var group models.Group
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(groupId)
		updateGroup(objId)

		err := groupCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": group}})
	}
}

//func UpdateGroup(id) / group name, genre, code

// func DeleteGroup(id)
func DeleteAGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupId := c.Param("groupId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(groupId)

		result, err := groupCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Group with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Group successfully deleted!"}},
		)
	}
}

// Mahian
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
								if likedMovie.MovieId.Id == objLikedMovieID {
									group.LikedMovies[idx].LikedCount++
									found = true
									break
								}
							}
							var placeHolderMovie models.Movie
							err := movieCollection.FindOne(ctx, bson.M{"id": objLikedMovieID}).Decode(&placeHolderMovie)
							if err != nil {
								log.Println("Error finding movie for ID:", objLikedMovieID)
								continue
							}
							if !found {
								group.LikedMovies = append(group.LikedMovies, models.Likedmovies{
									MovieId:    placeHolderMovie,
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
		objMovieID, _ := primitive.ObjectIDFromHex(movieID)
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
			if likedMovie.MovieId.Id == objMovieID {
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
		objMovieID, err2 := primitive.ObjectIDFromHex(movieID)
		objGroupID, err := primitive.ObjectIDFromHex(groupID)
		if err != nil || err2 != nil {
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
			if likedMovie.MovieId.Id == objMovieID {
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

		movieobjID, _ := primitive.ObjectIDFromHex(movieID)
		var placeHolderMovie models.Movie
		err1 := movieCollection.FindOne(ctx, bson.M{"id": movieobjID}).Decode(&placeHolderMovie)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": "Movie Doesn't Exist"},
			})
			return
		}
		group.LikedMovies = append(group.LikedMovies, models.Likedmovies{
			MovieId:    placeHolderMovie,
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
		updateGroup(objGroupID)

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
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": "No liked movies for user ID"},
			})
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
							if likedMovie.MovieId.Id == objLikedMovieID {
								group.LikedMovies[idx].LikedCount++
								found = true
								break
							}
						}
						var placeHolderMovie models.Movie
						err := movieCollection.FindOne(ctx, bson.M{"id": objLikedMovieID}).Decode(&placeHolderMovie)
						if err != nil {
							log.Println("Error finding movie for ID:", objLikedMovieID)
							continue
						}
						if !found {
							group.LikedMovies = append(group.LikedMovies, models.Likedmovies{
								MovieId:    placeHolderMovie,
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
