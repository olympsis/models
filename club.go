package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
Club
  - Club object
*/
type Club struct {
	ID           primitive.ObjectID  `json:"id,omitempty" bson:"_id"`
	Parent       Organization        `json:"parent,omitempty" bson:"parent, omitempty"`
	Name         string              `json:"name,omitempty" bson:"name"`
	Description  string              `json:"description,omitempty" bson:"description"`
	Sport        string              `json:"sport,omitempty" bson:"sport"`
	City         string              `json:"city,omitempty" bson:"city"`
	State        string              `json:"state,omitempty" bson:"state"`
	Country      string              `json:"country,omitempty" bson:"country"`
	ImageURL     string              `json:"image_url,omitempty" bson:"image_url"`
	ImageGallery []string            `json:"image_gallery,omitempty" bson:"image_gallery"`
	Visibility   string              `json:"visibility,omitempty" bson:"visibility"`
	Members      []Member            `json:"members,omitempty" bson:"members"`
	PinnedPostID *primitive.ObjectID `json:"pinned_post_id,omitempty" bson:"pinned_post_id"`
	Rules        []string            `json:"rules,omitempty" bson:"rules"`
	CreatedAt    int64               `json:"created_at,omitempty" bson:"created_at"`
}

type ClubResponse struct {
	Token string `json:"token,omitempty"`
	Club  Club   `json:"club"`
}

type ClubsResponse struct {
	TotalClubs int    `json:"total_clubs"`
	Clubs      []Club `json:"clubs"`
}

type ClubInvite struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	ClubID    string             `json:"club_id" bson:"club_id"`
	Data      *ClubInviteData    `json:"data,omitempty" bson:"data,omitempty"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

type ClubInviteData struct {
	Club    *Club     `json:"club,omitempty"`
	Inviter *UserData `json:"inviter,omitempty"`
}

type ClubInvitesResponse struct {
	TotalInvites int          `json:"total_invites"`
	Invites      []ClubInvite `json:"invites"`
}

type ChangeRoleRequest struct {
	Role string `json:"role"`
}

type CreateClubResponse struct {
	Token string `json:"token"`
	Club  Club   `json:"club"`
}

type ApplicationUpdateRequest struct {
	Status string `json:"status"`
}

type ClubApplication struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	ClubID    primitive.ObjectID `json:"club_id" bson:"club_id"`
	Status    string             `json:"status" bson:"status"`
	Data      *UserData          `json:"data,omitempty" bson:"data,omitempty"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

type ClubApplicationsResponse struct {
	TotalApplications int               `json:"total_applications"`
	Applications      []ClubApplication `json:"club_applications"`
}

type ClubDao struct {
	ParentID     *primitive.ObjectID `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	Name         *string             `json:"name,omitempty" bson:"name,omitempty"`
	Description  *string             `json:"description,omitempty" bson:"description,omitempty"`
	Sport        *string             `json:"sport,omitempty" bson:"sport,omitempty"`
	City         *string             `json:"city,omitempty" bson:"city,omitempty"`
	State        *string             `json:"state,omitempty" bson:"state,omitempty"`
	Country      *string             `json:"country,omitempty" bson:"country,omitempty"`
	ImageURL     *string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	ImageGallery *[]string           `json:"image_gallery" bson:"image_gallery,omitempty"`
	Visibility   *string             `json:"visibility,omitempty" bson:"visibility,omitempty"`
	Members      *[]MemberDao        `json:"members,omitempty" bson:"members,omitempty"`
	PinnedPostID *primitive.ObjectID `json:"pinned_post_id,omitempty" bson:"pinned_post_id,omitempty"`
	Rules        *[]string           `json:"rules,omitempty" bson:"rules,omitempty"`
	CreatedAt    *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
