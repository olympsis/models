package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Type            string             `json:"type" bson:"type"`
	Poster          UserSnippet        `json:"poster" bson:"poster"`
	Organizers      []Organizer        `json:"organizers" bson:"organizers"`
	Field           FieldDescriptor    `json:"field" bson:"field"`
	ImageURL        string             `json:"image_url" bson:"image_url"`
	Title           string             `json:"title" bson:"title"`
	Body            string             `json:"body" bson:"body"`
	Sport           string             `json:"sport" bson:"sport"`
	Level           int8               `json:"level" bson:"level"`
	StartTime       int64              `json:"start_time" bson:"start_time"`
	ActualStartTime int64              `json:"actual_start_time,omitempty" bson:"actual_start_time,omitempty"`
	StopTime        int64              `json:"stop_time" bson:"stop_time"`
	ActualStopTime  int64              `json:"actual_stop_time,omitempty" bson:"actual_stop_time,omitempty"`
	MinParticipants int64              `json:"min_participants" bson:"min_participants"`
	MaxParticipants int64              `json:"max_participants" bson:"max_participants"`
	Participants    []Participant      `json:"participants,omitempty" bson:"participants,omitempty"`
	Visibility      string             `json:"visibility" bson:"visibility"`
	Clubs           *[]Club            `json:"clubs,omitempty" bson:"clubs,omitempty"`
	Organizations   *[]Organization    `json:"organizations,omitempty" bson:"organizations,omitempty"`
	FieldData       *Field             `json:"field_data,omitempty" bson:"field_data,omitempty"`
	ExternalLink    string             `json:"external_link" bson:"external_link"`
	CreatedAt       int64              `json:"created_at" bson:"created_at"`
}

type EventsResponse struct {
	TotalEvents int     `json:"total_events"`
	Events      []Event `json:"events"`
}

type Participant struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at"`
}

type ParticipantDao struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt *int64             `json:"created_at,omitempty" bson:"created_at"`
}

type FullEventError struct {
	MSG   string `json:"msg"`
	Event Event  `json:"event"`
}

type Organizer struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

type FieldDescriptor struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type     string              `json:"type" bson:"type"`
	Name     *string             `json:"name,omitempty" bson:"name,omitempty"`
	Location *GeoJSON            `json:"location,omitempty" bson:"location,omitempty"`
}

type EventDao struct {
	Type            *string           `json:"type,omitempty" bson:"type,omitempty"`
	Poster          *string           `json:"poster,omitempty" bson:"poster,omitempty"`
	Organizers      *[]Organizer      `json:"organizers,omitempty" bson:"organizers,omitempty"`
	Field           *FieldDescriptor  `json:"field,omitempty" bson:"field,omitempty"`
	ImageURL        *string           `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Title           *string           `json:"title,omitempty" bson:"title,omitempty"`
	Body            *string           `json:"body,omitempty" bson:"body,omitempty"`
	Sport           *string           `json:"sport,omitempty" bson:"sport,omitempty"`
	Level           *int8             `json:"level,omitempty" bson:"level,omitempty"`
	StartTime       *int64            `json:"start_time,omitempty" bson:"start_time,omitempty"`
	ActualStartTime *int64            `json:"actual_start_time,omitempty" bson:"actual_start_time,omitempty"`
	StopTime        *int64            `json:"stop_time,omitempty" bson:"stop_time,omitempty"`
	ActualStopTime  *int64            `json:"actual_stop_time,omitempty" bson:"actual_stop_time,omitempty"`
	MinParticipants *int64            `json:"min_participants,omitempty" bson:"min_participants,omitempty"`
	MaxParticipants *int64            `json:"max_participants,omitempty" bson:"max_participants,omitempty"`
	Participants    *[]ParticipantDao `json:"participants,omitempty" bson:"participants,omitempty"`
	Visibility      *string           `json:"visibility,omitempty" bson:"visibility,omitempty"`
	ExternalLink    *string           `json:"external_link,omitempty" bson:"external_link,omitempty"`
	CreatedAt       *int64            `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
