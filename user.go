package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
User Data
  - Contains user identifiable data
*/
type User struct {
	ID            primitive.ObjectID    `json:"id,omitempty" bson:"_id"`
	UUID          string                `json:"uuid,omitempty" bson:"uuid"`
	UserName      string                `json:"username,omitempty" bson:"username"`
	Bio           string                `json:"bio,omitempty" bson:"bio,omitempty"`
	ImageURL      string                `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Visibility    string                `json:"visibility,omitempty" bson:"visibility"`
	Clubs         *[]primitive.ObjectID `json:"clubs,omitempty" bson:"clubs,omitempty"`
	Organizations *[]primitive.ObjectID `json:"organizations,omitempty" bson:"organizations,omitempty"`
	Sports        []string              `json:"sports,omitempty" bson:"sports,omitempty"`
	DeviceToken   string                `json:"device_token,omitempty" bson:"device_token,omitempty"`
}

// User data to return when looking up info about a user
type UserData struct {
	UUID          string                `json:"uuid" bson:"uuid"`
	Username      string                `json:"username" bson:"username"`
	FirstName     string                `json:"first_name" bson:"first_name"`
	LastName      string                `json:"last_name" bson:"last_name"`
	ImageURL      string                `json:"image_url" bson:"image_url"`
	Visibility    string                `json:"visibility" bson:"visibility"`
	Bio           string                `json:"bio,omitempty" bson:"bio,omitempty"`
	Clubs         *[]primitive.ObjectID `json:"clubs,omitempty" bson:"clubs,omitempty"`
	Organizations *[]primitive.ObjectID `json:"organizations,omitempty" bson:"organizations,omitempty"`
	Sports        []string              `json:"sports,omitempty" bson:"sports"`
	DeviceToken   string                `json:"device_token,omitempty" bson:"device_token,omitempty"`
}

// Users data response
type UsersDataResponse struct {
	TotalUsers int        `json:"total_users"`
	Users      []UserData `json:"users"`
}

type UserSnippet struct {
	UUID     string `json:"uuid" bson:"uuid"`
	Username string `json:"username" bson:"username"`
	ImageURL string `json:"image_url" bson:"image_url"`
}
