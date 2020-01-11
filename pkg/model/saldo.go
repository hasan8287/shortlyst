package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Saldo representing saldo coin
type Saldo struct {
	ObjectID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Value    string             `json:"value" bson:"value" validate:"required"`
	Count    int                `json:"count" bson:"count"`
}
