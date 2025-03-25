package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at"`
}
