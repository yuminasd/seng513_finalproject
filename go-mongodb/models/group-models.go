package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Likedmovies struct {
	MovieId    string `json:"movieId,omitempty"`
	LikedCount int    `json:"likedCount,omitempty"`
}

type Group struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"`
	Genre        []string           `json:"genre,omitempty"`
	Members      []User             `json:"members,omitempty"`
	LikedMovies  []Likedmovies      `json:"likedMovies,omitempty"`
}