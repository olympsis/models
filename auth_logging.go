package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID bson.ObjectID `bson:"user_id" json:"user_id"`

	SessionStart *bson.DateTime `bson:"session_start" json:"session_start"`
	SessionEnd   *bson.DateTime `bson:"session_end,omitempty" json:"session_end,omitempty"`
	LastActivity *bson.DateTime `bson:"last_activity" json:"last_activity"`

	DeviceInfo DeviceInfo `bson:"device_info" json:"device_info"`

	CreatedAt bson.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt bson.DateTime `bson:"updated_at" json:"updated_at"`
}
