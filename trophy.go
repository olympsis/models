package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
Trophy
  - Trophy object
*/
type Trophy struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Title       string             `json:"title" bson:"title"`
	ImageURL    string             `json:"image_url" bson:"image_url"`
	EventId     primitive.ObjectID `json:"event_id" bson:"event_id"`
	Description string             `json:"description" bson:"description"`
	AchievedAt  primitive.DateTime `json:"achieved_at,omitempty" bson:"achieved_at,omitempty"`
}
