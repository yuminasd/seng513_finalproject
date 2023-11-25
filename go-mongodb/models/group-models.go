package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	Id       	primitive.ObjectID 	`json:"id,omitempty"`
	Name     	string             	`json:"name,omitempty" validate:"required"`
	Code	 	string				`json:"code,omitempty"`
	Genre 	 	[]string			`json:"genre,omitempty"`
	Members 	[]User				`json:"User,omitempty"`
	LikedMovies []string			`json:"likedMovies,omitempty"`
}
