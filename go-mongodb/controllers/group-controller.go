package controllers

import (
	"go-mongodb/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var groupCollection *mongo.Collection = configs.GetCollection(configs.DB, "groups")
var validate = validator.New()

//Zainab
//func CreateGroup()
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
		if validationErr := validate.Struct(&group); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newGroup := models.Group{
			Id:           primitive.NewObjectID(),
			Name:         group.Name,
			Code:		  group.Code,
			Genre:        group.Genre,
			Members:      group.Members,
			LikedMovies:  group.LikedMovies,
		}

		result, err := groupCollection.InsertOne(ctx, newGroup)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

//func GetGroupInfo(id) / get all name, users, genre
func GetGroupInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		groupId := c.Param("groupId")
		var group models.Group
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(groupId)

		err := groupCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": group}})
	}
}

//func UpdateGroup(id) / group name, genre, code

//func DeleteGroup(id)
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

//Mahian1
//func GetLikedMovies(id) / returning movies most liked by users

//func RemoveLikedMovies(Movie.Id)

//func RemoveUser(User.id)


//GetAllGroups for User Controller?