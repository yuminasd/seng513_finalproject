package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Name           string             `json:"name,omitempty" validate:"required"`
	EmailAddress   string             `json:"emailaddress,omitempty" validate:"required"`
	Password       string             `json:"password,omitempty" validate:"required"`
	Image          string             `json:"image,omitempty" validate:"required"`
	GroupID        []string           `json:"groupID,omitempty"`
	LikedMovies    []string           `json:"likedMovies,omitempty"`
	DislikedMovies []string           `json:"dislikedMovies,omitempty"`
}

/*
Cannot think of anything else we would need besides movieID so made it a string can change

type LikeDislikeUser struct {
	MovieID string
}
*/
