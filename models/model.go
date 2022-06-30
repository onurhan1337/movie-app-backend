package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id         primitive.ObjectID `json:"_id,omitempty"`
	Title      string             `json:"title,omitempty" validate:"required"`
	ImdbRating float64            `json:"imdbRating,omitempty" validate:"required"`
	Image      string             `json:"image,omitempty" validate:"required"`
	Year       int                `json:"year,omitempty" validate:"required"`
	Director   string             `json:"director,omitempty" validate:"required"`
}
