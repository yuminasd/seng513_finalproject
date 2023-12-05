package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Img         string             `json:"img,omitempty" validate:"required"`
	BgImg       string             `json:"bgImg,omitempty" validate:"required"`
	Rating      int                `json:"rating,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	Genres      []string           `json:"genres,omitempty" validate:"required"`
}
