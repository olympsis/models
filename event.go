package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Type            int8               `json:"type" bson:"type"`
	Poster          UserSnippet        `json:"poster" bson:"poster"`
	Organizers      []Organizer        `json:"organizers" bson:"organizers"`
	Venues          []VenueDescriptor  `json:"venues" bson:"venues"`
	ImageURL        string             `json:"image_url" bson:"image_url"`
	Title           string             `json:"title" bson:"title"`
	Body            string             `json:"body" bson:"body"`
	Sports          []string           `json:"sports" bson:"sports"`
	Level           int8               `json:"level" bson:"level"`
	StartTime       int64              `json:"start_time" bson:"start_time"`
	StopTime        int64              `json:"stop_time" bson:"stop_time"`
	MinParticipants *int32             `json:"min_participants,omitempty" bson:"min_participants,omitempty"`
	MaxParticipants *int32             `json:"max_participants,omitempty" bson:"max_participants,omitempty"`
	Participants    *[]Participant     `json:"participants,omitempty" bson:"participants,omitempty"`
	Visibility      int8               `json:"visibility" bson:"visibility"`
	ExternalLink    *string            `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive     bool               `json:"is_sensitive" bson:"is_sensitive"`
	CreatedAt       int64              `json:"created_at" bson:"created_at"`
}

type EventDao struct {
	Type            *int8              `json:"type,omitempty" bson:"type,omitempty"`
	Poster          *string            `json:"poster,omitempty" bson:"poster,omitempty"`
	Organizers      *[]Organizer       `json:"organizers,omitempty" bson:"organizers,omitempty"`
	Venues          *[]VenueDescriptor `json:"venues,omitempty" bson:"venues,omitempty"`
	ImageURL        *string            `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Title           *string            `json:"title,omitempty" bson:"title,omitempty"`
	Body            *string            `json:"body,omitempty" bson:"body,omitempty"`
	Sports          *[]string          `json:"sports,omitempty" bson:"sports,omitempty"`
	Level           *int8              `json:"level,omitempty" bson:"level,omitempty"`
	StartTime       *int64             `json:"start_time,omitempty" bson:"start_time,omitempty"`
	StopTime        *int64             `json:"stop_time,omitempty" bson:"stop_time,omitempty"`
	MinParticipants *int32             `json:"min_participants,omitempty" bson:"min_participants,omitempty"`
	MaxParticipants *int32             `json:"max_participants,omitempty" bson:"max_participants,omitempty"`
	Participants    *[]ParticipantDao  `json:"participants,omitempty" bson:"participants,omitempty"`
	Visibility      *int8              `json:"visibility,omitempty" bson:"visibility,omitempty"`
	ExternalLink    *string            `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive     *bool              `json:"is_sensitive,omitempty" bson:"is_sensitive,omitempty"`
	CreatedAt       *int64             `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type EventsResponse struct {
	TotalEvents int16   `json:"total_events"`
	Events      []Event `json:"events"`
}

type Participant struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Status    int8               `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at"`
}

type ParticipantDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UUID      *string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Status    *int8               `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type FullEventError struct {
	MSG   string `json:"msg"`
	Event Event  `json:"event"`
}

type Organizer struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type int8               `json:"type" bson:"type"`
}
