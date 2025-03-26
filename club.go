package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Club struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Parent      *OrganizationDao     `json:"parent,omitempty" bson:"parent,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Tags        []string             `json:"tags" bson:"tags"`
	Sports      []string             `json:"sports" bson:"sports"`
	City        string               `json:"city" bson:"city"`
	State       string               `json:"state" bson:"state"`
	Country     string               `json:"country" bson:"country"`
	Logo        string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      string               `json:"banner,omitempty" bson:"banner,omitempty"`
	Visibility  string               `json:"visibility" bson:"visibility"`
	Members     []Member             `json:"members" bson:"members"`
	PinnedPosts []primitive.ObjectID `json:"pinned_posts" bson:"pinned_posts"`
	BlackList   []string             `json:"black_list" bson:"black_list"`
	Rules       []string             `json:"rules" bson:"rules"`
	SnapshotURL *string              `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
	IsVerified  bool                 `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	CreatedAt   primitive.DateTime   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type ClubDao struct {
	ID          *primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	ParentID    *primitive.ObjectID   `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	Name        *string               `json:"name,omitempty" bson:"name,omitempty"`
	Description *string               `json:"description,omitempty" bson:"description,omitempty"`
	Tags        []string              `json:"tags,omitempty" bson:"tags,omitempty"`
	Sports      *[]string             `json:"sports,omitempty" bson:"sports,omitempty"`
	City        *string               `json:"city,omitempty" bson:"city,omitempty"`
	State       *string               `json:"state,omitempty" bson:"state,omitempty"`
	Country     *string               `json:"country,omitempty" bson:"country,omitempty"`
	Logo        *string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      *string               `json:"banner,omitempty" bson:"banner,omitempty"`
	Visibility  *string               `json:"visibility,omitempty" bson:"visibility,omitempty"`
	PinnedPosts *[]primitive.ObjectID `json:"pinned_posts,omitempty" bson:"pinned_posts,omitempty"`
	BlackList   *[]string             `json:"black_list,omitempty" bson:"black_list,omitempty"`
	Rules       *[]string             `json:"rules,omitempty" bson:"rules,omitempty"`
	SnapshotURL *string               `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
	IsVerified  *bool                 `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	CreatedAt   *primitive.DateTime   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for club objects
type ClubsResponse struct {
	TotalClubs int    `json:"total_clubs"`
	Clubs      []Club `json:"clubs"`
}

// Object to handle club invitations
type ClubInvite struct {
	ID        primitive.ObjectID `json:"id"`
	Club      *Club              `json:"club,omitempty"`
	Sender    *UserData          `json:"sender,omitempty"`
	Status    string             `json:"status"`
	CreatedAt primitive.DateTime `json:"created_at"`
}

// Data access object for club invites
type ClubInviteDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Recipient *string             `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Sender    *string             `json:"sender,omitempty" bson:"sender,omitempty"`
	ClubID    *string             `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for the club invite objects
type ClubInvitesResponse struct {
	TotalInvites int          `json:"total_invites"`
	Invites      []ClubInvite `json:"invites"`
}

// Object to handle club applications
type ClubApplication struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Applicant *UserData          `json:"applicant,omitempty"`
	Status    string             `json:"status"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

// Data access object for club applications
type ClubApplicationDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Applicant *string             `json:"applicant,omitempty" bson:"applicant,omitempty"`
	ClubID    *primitive.ObjectID `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for club applications
type ClubApplicationsResponse struct {
	TotalApplications int               `json:"total_applications"`
	Applications      []ClubApplication `json:"club_applications"`
}

// Object to handle role change requests
type ChangeRoleRequest struct {
	Role string `json:"role"`
}
