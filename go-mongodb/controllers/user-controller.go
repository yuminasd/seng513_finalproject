package controllers

import (
	"context"
	"go-mongodb/configs"
	"go-mongodb/models"
	"go-mongodb/responses"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admin")

var validate = validator.New()

func updateUserGroups(userID primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		return
	}

	var newGroupList []models.Group
	var newGroup models.Group
	for i := 0; i < len(user.GroupID); i++ {
		groupID := user.GroupID[i].Id
		err := groupCollection.FindOne(ctx, bson.M{"id": groupID}).Decode(&newGroup)
		if err != nil {
			continue
		}
		newGroup.Members = nil
		newGroupList = append(newGroupList, newGroup)
	}

	user.GroupID = newGroupList
	userCollection.ReplaceOne(ctx, bson.M{"id": userID}, user)
}

func indexOf(element string, data []models.Group) int {
	for i := 0; i < len(data); i++ {
		log.Printf(data[i].Id.String(), " \n", element)
		if strings.Contains(data[i].Id.String(), element) {
			return i
		}
	}
	return -1
}

func indexOfString(element string, data []string) int {
	for i := 0; i < len(data); i++ {
		if data[i] == element {
			return i
		}
	}
	return -1
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"emailaddress": string(user.EmailAddress)}).Decode(&user)
		if err == nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "User already exists!"}})
			return
		}

		newUser := models.User{
			Id:           primitive.NewObjectID(),
			Name:         user.Name,
			EmailAddress: user.EmailAddress,
			Password:     user.Password,
			Image:        user.Image,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		updateUserGroups(objId)
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			updateUserGroups(singleUser.Id)
		}

		result, err := userCollection.Find(ctx, bson.M{})

		//Somewhat redundant code but has to loop through list of users to update group to make sure groups are up to date
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer result.Close(ctx)
		for result.Next(ctx) {
			var singleUser models.User
			if err = result.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()
		var user models.User

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		_, err := userCollection.UpdateByID(ctx, userId, UpdatedUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)
	}
}

func AddGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		groupId := c.Param("groupId")
		defer cancel()
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		groupID, _ := primitive.ObjectIDFromHex(groupId)
		var newGroup models.Group
		err2 := groupCollection.FindOne(ctx, bson.M{"id": groupID}).Decode(&newGroup)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err2.Error()}})
			return
		}

		if user.GroupID == nil {
			var newGroupList []models.Group
			newGroupList = append(newGroupList, newGroup)
			UpdatedUser.GroupID = newGroupList
		} else {
			UpdatedUser.GroupID = append(UpdatedUser.GroupID, newGroup)
		}

		_, err1 := userCollection.ReplaceOne(ctx, bson.M{"id": objId}, UpdatedUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err1.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)

	}
}

func RemoveGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		groupId := c.Param("groupId")
		defer cancel()
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		var index = indexOf(groupId, user.GroupID)
		if user.GroupID != nil && index != -1 {
			if len(user.GroupID) == 1 {
				UpdatedUser.GroupID = nil

			} else {
				UpdatedUser.GroupID = slices.Delete(UpdatedUser.GroupID, index, index+1)
			}
		}

		_, err1 := userCollection.ReplaceOne(ctx, bson.M{"id": objId}, UpdatedUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err1.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)
	}
}

func AddLiked() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		movieId := c.Param("movieId")
		defer cancel()
		var user models.User
		var movie models.Movie
		objId, _ := primitive.ObjectIDFromHex(userId)
		movieID, _ := primitive.ObjectIDFromHex(movieId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		err2 := movieCollection.FindOne(ctx, bson.M{"id": movieID}).Decode(&movie)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err2.Error()}})
			return
		}

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		if user.LikedMovies == nil {
			var newLiked []string
			newLiked = append(newLiked, movieId)
			UpdatedUser.LikedMovies = newLiked
		} else {
			UpdatedUser.LikedMovies = append(UpdatedUser.LikedMovies, movieId)
		}

		_, err1 := userCollection.ReplaceOne(ctx, bson.M{"id": objId}, UpdatedUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err1.Error()}})
			return
		}

		for i := 0; i < len(UpdatedUser.GroupID); i++ {
			for k := 0; k < len(movie.Genres); k++ {
				if slices.Contains(UpdatedUser.GroupID[i].Genre, movie.Genres[k]) {
					AddMovieToGroupLiked(movieID, UpdatedUser.GroupID[i].Id, c, ctx)
					break
				}
			}
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)
	}
}

func AddMovieToGroupLiked(objMovieID primitive.ObjectID, objGroupID primitive.ObjectID, c *gin.Context, ctx context.Context) {
	var group models.Group
	var err = groupCollection.FindOne(ctx, bson.M{"id": objGroupID}).Decode(&group)
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
			return
		}
	}

	var placeHolderMovie models.Movie
	err1 := movieCollection.FindOne(ctx, bson.M{"id": objMovieID}).Decode(&placeHolderMovie)
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
}

func RemoveLiked() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		movieId := c.Param("movieId")
		defer cancel()
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		var index = indexOfString(movieId, UpdatedUser.LikedMovies)
		if user.LikedMovies != nil && index != -1 {
			UpdatedUser.LikedMovies = slices.Delete(UpdatedUser.LikedMovies, index, index+1)
		}

		_, err1 := userCollection.ReplaceOne(ctx, bson.M{"id": objId}, UpdatedUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err1.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)
	}
}

func AddDisliked() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		movieId := c.Param("movieId")
		defer cancel()
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		if user.LikedMovies == nil {
			var newLiked []string
			newLiked = append(newLiked, movieId)
			UpdatedUser.DislikedMovies = newLiked
		} else {
			UpdatedUser.DislikedMovies = append(UpdatedUser.DislikedMovies, movieId)
		}

		_, err1 := userCollection.ReplaceOne(ctx, bson.M{"id": objId}, UpdatedUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err1.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)
	}
}

func RemoveDisliked() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		movieId := c.Param("movieId")
		defer cancel()
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		UpdatedUser := models.User{
			Id:             user.Id,
			Name:           user.Name,
			EmailAddress:   user.EmailAddress,
			Password:       user.Password,
			Image:          user.Image,
			GroupID:        user.GroupID,
			LikedMovies:    user.LikedMovies,
			DislikedMovies: user.DislikedMovies,
		}

		var index = indexOfString(movieId, UpdatedUser.DislikedMovies)
		if user.DislikedMovies != nil && index != -1 {
			UpdatedUser.DislikedMovies = slices.Delete(UpdatedUser.DislikedMovies, index, index+1)
		}

		_, err1 := userCollection.ReplaceOne(ctx, bson.M{"id": objId}, UpdatedUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err1.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully updated!"}},
		)
	}
}

func GetLiked() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var movieList []models.Movie
		var movie models.Movie
		for i := 0; i < len(user.LikedMovies); i++ {
			var movieID, _ = primitive.ObjectIDFromHex(user.LikedMovies[i])
			err := movieCollection.FindOne(ctx, bson.M{"id": movieID}).Decode(&movie)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			movieList = append(movieList, movie)
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movieList}})
	}
}

func GetDisliked() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user.DislikedMovies}})
	}
}

type LoginRequestBody struct {
	Email    string
	Password string
}

type Admin struct {
	Id primitive.ObjectID `json:"id,omitempty"`
}

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var requestBody LoginRequestBody
		var user models.User
		defer cancel()

		if errbody := c.BindJSON(&requestBody); errbody != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errbody.Error()}})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"emailaddress": requestBody.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var admin Admin
		var role string
		err1 := adminCollection.FindOne(ctx, bson.M{"id": user.Id}).Decode(&admin)
		if err1 != nil {
			role = "user"
		} else {
			role = "admin"
		}

		if user.Password != requestBody.Password {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Not Authenticated"}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Authentication Successful", "id": user.Id, "userRole": role}})
	}
}
