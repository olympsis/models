package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Announcement struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Body      string             `json:"body" bson:"body"`
	MediaURL  string             `json:"media_url" bson:"media_url"`
	Action    string             `json:"action,omitempty" bson:"action,omitempty"`
	ActionURL string             `json:"action_url,omitempty" bson:"action_url,omitempty"`
	StartTime int64              `json:"start_time" bson:"start_time"`
	EndTime   *int64             `json:"end_time,omitempty" bson:"end_time,omitempty"`
	Location  *GeoJSON           `json:"location,omitempty" bson:"location,omitempty"`
	Country   string             `json:"country,omitempty" bson:"country,omitempty"`
	State     string             `json:"state,omitempty" bson:"state,omitempty"`
	City      string             `json:"city,omitempty" bson:"city,omitempty"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
}

type AnnouncementDao struct {
	Title     *string  `json:"title,omitempty" bson:"title,omitempty"`
	Body      *string  `json:"body,omitempty" bson:"body,omitempty"`
	MediaURL  *string  `json:"media_url,omitempty" bson:"media_url,omitempty"`
	Action    *string  `json:"action,omitempty" bson:"action,omitempty"`
	ActionURL *string  `json:"action_url,omitempty" bson:"action_url,omitempty"`
	StartTime *int64   `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime   *int64   `json:"end_time,omitempty" bson:"end_time,omitempty"`
	Location  *GeoJSON `json:"location,omitempty" bson:"location,omitempty"`
	Country   *string  `json:"country,omitempty" bson:"country,omitempty"`
	State     *string  `json:"state,omitempty" bson:"state,omitempty"`
	City      *string  `json:"city,omitempty" bson:"city,omitempty"`
	CreatedAt *int64   `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy *string  `json:"created_by,omitempty" bson:"created_by,omitempty"`
}

type AnnouncementsResponse struct {
	TotalAnnouncements int16          `json:"total_announcements"`
	Announcements      []Announcement `json:"announcements"`
}
