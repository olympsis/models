package models

import "go.mongodb.org/mongo-driver/v2/bson"

/*
Trophy
  - Trophy object
*/
type Trophy struct {
	ID          bson.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Title       string             `json:"title" bson:"title"`
	ImageURL    string             `json:"image_url" bson:"image_url"`
	EventId     bson.ObjectID `json:"event_id" bson:"event_id"`
	Description string             `json:"description" bson:"description"`
	AchievedAt  bson.DateTime `json:"achieved_at,omitempty" bson:"achieved_at,omitempty"`
}
