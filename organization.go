package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organization struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Description  string             `json:"description" bson:"description"`
	Sport        string             `json:"sport" bson:"sport"`
	City         string             `json:"city" bson:"city"`
	State        string             `json:"state" bson:"state"`
	Country      string             `json:"country" bson:"country"`
	ImageURL     string             `json:"image_url" bson:"image_url"`
	ImageGallery []string           `json:"image_gallery" bson:"image_gallery"`
	Members      []Member           `json:"members" bson:"members"`
	PinnedPostID primitive.ObjectID `json:"pinned_post_id" bson:"pinned_post_id"`
	Data         *OrganizationData  `json:"data" bson:"data,omitempty"`
	CreatedAt    int64              `json:"created_at" bson:"created_at"`
}

type OrganizationData struct {
	Children *[]Club `json:"children"`
}

type OrganizationsResponse struct {
	TotalEvents   int            `json:"total_organizations"`
	Organizations []Organization `json:"organizations"`
}
