package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organization struct {
	ID           primitive.ObjectID  `json:"id" bson:"_id"`
	Name         string              `json:"name" bson:"name"`
	Description  string              `json:"description" bson:"description"`
	Sport        string              `json:"sport" bson:"sport"`
	City         string              `json:"city" bson:"city"`
	State        string              `json:"state" bson:"state"`
	Country      string              `json:"country" bson:"country"`
	ImageURL     string              `json:"image_url" bson:"image_url"`
	ImageGallery []string            `json:"image_gallery" bson:"image_gallery"`
	Members      []Member            `json:"members" bson:"members"`
	Children     *[]ClubDao          `json:"children,omitempty" bson:"children,omitempty"`
	PinnedPostID *primitive.ObjectID `json:"pinned_post_id,omitempty" bson:"pinned_post_id,omitempty"`
	CreatedAt    int64               `json:"created_at" bson:"created_at"`
}

type OrganizationsResponse struct {
	TotalEvents   int            `json:"total_organizations"`
	Organizations []Organization `json:"organizations"`
}

type OrganizationApplication struct {
	ID             primitive.ObjectID           `json:"id" bson:"_id"`
	ClubID         primitive.ObjectID           `json:"club_id" bson:"club_id"`
	OrganizationID primitive.ObjectID           `json:"organization_id" bson:"organization_id"`
	Status         string                       `json:"status" bson:"status"`
	Data           *OrganizationApplicationData `json:"data,omitempty" bson:"data,omitempty"`
	CreatedAt      int64                        `json:"created_at" bson:"created_at"`
}

type OrganizationApplicationData struct {
	Club         *Club         `json:"club,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

type OrganizationDao struct {
	ID           *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         *string             `json:"name,omitempty" bson:"name,omitempty"`
	Description  *string             `json:"description,omitempty" bson:"description,omitempty"`
	Sport        *string             `json:"sport,omitempty" bson:"sport,omitempty"`
	City         *string             `json:"city,omitempty" bson:"city,omitempty"`
	State        *string             `json:"state,omitempty" bson:"state,omitempty"`
	Country      *string             `json:"country,omitempty" bson:"country,omitempty"`
	ImageURL     *string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	ImageGallery *[]string           `json:"image_gallery" bson:"image_gallery,omitempty"`
	Members      *[]MemberDao        `json:"members,omitempty" bson:"members,omitempty"`
	PinnedPostID *primitive.ObjectID `json:"pinned_post_id,omitempty" bson:"pinned_post_id,omitempty"`
	CreatedAt    *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
