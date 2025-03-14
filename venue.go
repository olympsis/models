package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Venue struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Owner       Ownership          `json:"owner" bson:"owner"`
	Description string             `json:"description" bson:"description"`
	Sports      []string           `json:"sports" bson:"sports"`
	Images      []string           `json:"images" bson:"images"`
	City        string             `json:"city" bson:"city"`
	State       string             `json:"state" bson:"state"`
	Country     string             `json:"country" bson:"country"`
	Location    GeoJSON            `json:"location" bson:"location"`
	SnapshotURL *string            `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
}

type Ownership struct {
	Name    string              `json:"name" bson:"name"`
	Type    string              `json:"type" bson:"type"`
	OwnerID *primitive.ObjectID `json:"owner_id,omitempty" bson:"owner_id,omitempty"`
}

type VenuesResponse struct {
	TotalVenues int     `json:"total_venues"`
	Venues      []Venue `json:"venues"`
}

type VenueDescriptor struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     *string             `json:"name,omitempty" bson:"name,omitempty"`
	City     *string             `json:"city,omitempty" bson:"city,omitempty"`
	State    *string             `json:"state,omitempty" bson:"state,omitempty"`
	Country  *string             `json:"country,omitempty" bson:"country,omitempty"`
	Location *GeoJSON            `json:"location,omitempty" bson:"location,omitempty"`
}
