package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Type            string             `json:"type" bson:"type"`
	Poster          string             `json:"poster" bson:"poster"`
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
	Data            *EventData         `json:"data" bson:"data,omitempty"`
	ExternalLink    string             `json:"external_link" bson:"external_link"`
	CreatedAt       int64              `json:"created_at" bson:"created_at"`
}

type EventData struct {
	Poster        *UserData       `json:"poster,omitempty"`
	Clubs         *[]Club         `json:"clubs,omitempty"`
	Organizations *[]Organization `json:"organizations,omitempty"`
	Field         *Field          `json:"field,omitempty"`
}

type EventsResponse struct {
	TotalEvents int     `json:"total_events"`
	Events      []Event `json:"events"`
}

type Participant struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Data      *UserData          `json:"data,omitempty" bson:"data,omitempty"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at"`
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
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Type     string             `json:"type" bson:"type"`
	Location *GeoJSON           `json:"location,omitempty" bson:"location,omitempty"`
}
