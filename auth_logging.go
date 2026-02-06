package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`

	SessionStart *primitive.DateTime `bson:"session_start" json:"session_start"`
	SessionEnd   *primitive.DateTime `bson:"session_end,omitempty" json:"session_end,omitempty"`
	LastActivity *primitive.DateTime `bson:"last_activity" json:"last_activity"`

	DeviceInfo DeviceInfo `bson:"device_info" json:"device_info"`

	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}
