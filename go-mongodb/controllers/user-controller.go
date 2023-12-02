package controllers

import (
	"context"
	"go-mongodb/configs"
	"go-mongodb/models"
	"go-mongodb/responses"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

var validate = validator.New()

func indexOf(element string, data []string) int {
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

		if user.GroupID == nil {
			var newGroup []string
			newGroup = append(newGroup, groupId)
			UpdatedUser.GroupID = newGroup
		} else {
			UpdatedUser.GroupID = append(UpdatedUser.GroupID, groupId)
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

		var index = indexOf(groupId, UpdatedUser.GroupID)
		if user.GroupID != nil && index != -1 {
			UpdatedUser.GroupID = slices.Delete(UpdatedUser.GroupID, index, index+1)
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
			UpdatedUser.LikedMovies = newLiked
		} else {
			UpdatedUser.LikedMovies = append(UpdatedUser.LikedMovies, movieId)
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

		var index = indexOf(movieId, UpdatedUser.LikedMovies)
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

		var index = indexOf(movieId, UpdatedUser.DislikedMovies)
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

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user.LikedMovies}})
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

		if user.Password != requestBody.Password {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Not Authenticated"}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Authentication Successful", "id": user.Id}})
	}
}
