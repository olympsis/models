package models

import "go.mongodb.org/mongo-driver/v2/bson"

// Object represents groups that clubs can organize under for a hierarchy
type Organization struct {
	ID          bson.ObjectID   `json:"id" bson:"_id"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Tags        []string             `json:"tags" bson:"tags"`
	Sports      []string             `json:"sports" bson:"sports"`
	City        string               `json:"city" bson:"city,"`
	State       string               `json:"state" bson:"state"`
	Country     string               `json:"country" bson:"country"`
	Location    GeoJSON              `json:"location" bson:"location"`
	Logo        string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      string               `json:"banner,omitempty" bson:"banner,omitempty"`
	Member      []Member             `json:"members" bson:"members"`
	Children    []bson.ObjectID `json:"children" bson:"children"`
	BlackList   []string             `json:"black_list" bson:"black_list"`
	PinnedPosts []bson.ObjectID `json:"pinned_posts" bson:"pinned_posts"`
	SnapshotURL *string              `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
	IsVerified  bool                 `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	CreatedAt   bson.DateTime   `json:"created_at" bson:"created_at"`
}

// Data access object for the organizations
type OrganizationDao struct {
	ID          *bson.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        *string               `json:"name,omitempty" bson:"name,omitempty"`
	Description *string               `json:"description,omitempty" bson:"description,omitempty"`
	Tags        []string              `json:"tags,omitempty" bson:"tags,omitempty"`
	Sports      *[]string             `json:"sports,omitempty" bson:"sports,omitempty"`
	City        *string               `json:"city,omitempty" bson:"city,omitempty"`
	State       *string               `json:"state,omitempty" bson:"state,omitempty"`
	Country     *string               `json:"country,omitempty" bson:"country,omitempty"`
	Location    GeoJSON               `json:"location,omitempty" bson:"location,omitempty"`
	Logo        *string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      *string               `json:"banner,omitempty" bson:"banner,omitempty"`
	BlackList   *[]string             `json:"black_list,omitempty" bson:"black_list,omitempty"`
	PinnedPosts *[]bson.ObjectID `json:"pinned_posts,omitempty" bson:"pinned_posts,omitempty"`
	SnapshotURL *string               `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
	IsVerified  *bool                 `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	CreatedAt   *bson.DateTime   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for organization objects
type OrganizationsResponse struct {
	TotalOrganizations int            `json:"total_organizations"`
	Organizations      []Organization `json:"organizations"`
}

// Object to handle organization applications
type OrganizationApplication struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	Status    string             `json:"status" bson:"status"`
	Club      *ClubDao           `json:"club,omitempty" bson:"club,omitempty"`
	CreatedAt bson.DateTime `json:"created_at" bson:"created_at"`
}

// Data access object for organizations applications
type OrganizationApplicationDao struct {
	ID             *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ClubID         *bson.ObjectID `json:"club_id,omitempty" bson:"club_id,omitempty"`
	OrganizationID *bson.ObjectID `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	Status         *string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt      *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for organization applications
type OrganizationApplicationsResponse struct {
	TotalApplications int                       `json:"total_applications"`
	Applications      []OrganizationApplication `json:"organization_applications"`
}
