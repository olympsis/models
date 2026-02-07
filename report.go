package models

import "go.mongodb.org/mongo-driver/v2/bson"

type PostReport struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	Post      *Post               `json:"post,omitempty" bson:"post,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type PostReportDao struct {
	ID        *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PostID    *bson.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	GroupID   *bson.ObjectID `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type MemberReport struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	Member    *UserSnippet        `json:"member,omitempty" bson:"member,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type MemberReportDao struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	MemberID  *bson.ObjectID `json:"member_id,omitempty" bson:"member_id,omitempty"`
	GroupID   *bson.ObjectID `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type EventReport struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	User      *UserData           `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Event     *Event              `json:"event,omitempty" bson:"event,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type EventReportDao struct {
	ID        *bson.ObjectID   `json:"id" bson:"_id"`
	User      *string               `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string               `json:"type,omitempty" bson:"type,omitempty"`
	EventID   *bson.ObjectID   `json:"event_id,omitempty" bson:"event_id,omitempty"`
	Groups    *[]bson.ObjectID `json:"groups,omitempty" bson:"groups,omitempty"`
	Notes     *string               `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string               `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message            `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type VenueReport struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	User      *UserData           `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	Venue     *Venue              `json:"venue,omitempty" bson:"venue,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type VenueReportDao struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	User      *string             `json:"user,omitempty" bson:"user,omitempty"`
	Type      *string             `json:"type,omitempty" bson:"type,omitempty"`
	VenueID   *bson.ObjectID `json:"venue_id,omitempty" bson:"venue_id,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type BugReport struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	User      *UserData           `json:"user,omitempty" bson:"user,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Images    *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	Videos    *[]string           `json:"videos,omitempty" bson:"videos,omitempty"`
	Blobs     *[]string           `json:"blobs,omitempty" bson:"blobs,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type BugReportDao struct {
	ID        *bson.ObjectID `json:"id" bson:"_id"`
	User      *string             `json:"user,omitempty" bson:"user,omitempty"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	Images    *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	Videos    *[]string           `json:"videos,omitempty" bson:"videos,omitempty"`
	Blobs     *[]string           `json:"blobs,omitempty" bson:"blobs,omitempty"`
	Messages  *[]Message          `json:"messages,omitempty" bson:"messages,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
