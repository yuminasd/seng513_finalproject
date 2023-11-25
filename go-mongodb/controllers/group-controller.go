package controllers

import (
	"go-mongodb/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var groupCollection *mongo.Collection = configs.GetCollection(configs.DB, "groups")

//Zainab
//func CreateGroup

//func GetGroupInfo(id) / get all name, users, genre

//func UpdateGroup(id) / group name, genre, code

//func DeleteGroup(id)

//Mahian1
//func GetLikedMovies(id) / returning movies most liked by users

//func RemoveLikedMovies(Movie.Id)

//func RemoveUser(User.id)
