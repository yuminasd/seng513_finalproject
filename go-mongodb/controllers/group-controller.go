package controllers

import (
	"go-mongodb/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var groupCollection *mongo.Collection = configs.GetCollection(configs.DB, "groups")
