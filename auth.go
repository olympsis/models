package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Authentication User Data
type AuthUser struct {
	UUID      string              `json:"uuid" bson:"uuid"`
	FirstName string              `json:"first_name" bson:"first_name"`
	LastName  string              `json:"last_name" bson:"last_name"`
	Email     string              `json:"email" bson:"email"`
	Gender    *Gender             `json:"gender,omitempty" bson:"gender,omitempty"`
	Birthdate *primitive.DateTime `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	CreatedAt primitive.DateTime  `json:"created_at" bson:"created_at"`
}

type AuthUserDao struct {
	UUID      *string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	FirstName *string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  *string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email     *string             `json:"email,omitempty" bson:"email,omitempty"`
	Gender    *Gender             `json:"gender,omitempty" bson:"gender,omitempty"`
	Birthdate *primitive.DateTime `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
}

// Authentication Request
type AuthRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Token     *string `json:"token,omitempty"`
	UUID      *string `json:"uuid,omitempty"`
}

// Authentication Response
type AuthResponse struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
