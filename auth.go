package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Authentication User Data
type AuthUser struct {
	UUID      *string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	FirstName *string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  *string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email     *string             `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
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
	UUID      *string `json:"uuid,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
}
