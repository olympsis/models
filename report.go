package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostReport struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	Post      *Post               `json:"post,omitempty" bson:"post,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type PostReportDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PostID    *primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	GroupID   *primitive.ObjectID `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type MemberReport struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	Member    *Member             `json:"member,omitempty" bson:"member,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type MemberReportDao struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	MemberID  *primitive.ObjectID `json:"member_id,omitempty" bson:"member_id,omitempty"`
	GroupID   *primitive.ObjectID `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type EventReport struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	User      *UserData           `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Event     *Event              `json:"event,omitempty" bson:"event,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type EventReportDao struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	User      *string             `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	EventID   *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type FieldReport struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	User      *UserData           `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Field     *Field              `json:"field,omitempty" bson:"field,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type FieldReportDao struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	User      *string             `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	FieldID   *primitive.ObjectID `json:"field_id,omitempty" bson:"field_id,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type BugReport struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	User      *UserData           `json:"user,omitempty" bson:"user,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Images    *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	Videos    *[]string           `json:"videos,omitempty" bson:"videos,omitempty"`
	Blobs     *[]string           `json:"blobs,omitempty" bson:"blobs,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type BugReportDao struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	User      *string             `json:"user,omitempty" bson:"user,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Images    *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	Videos    *[]string           `json:"videos,omitempty" bson:"videos,omitempty"`
	Blobs     *[]string           `json:"blobs,omitempty" bson:"blobs,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
