package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Items representing item menu
type Items struct {
	ObjectID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Price    int64              `json:"price" bson:"price" validate:"required"`
	Stock    int                `json:"stock" bson:"stock"`
}
