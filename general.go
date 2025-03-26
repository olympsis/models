package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Invitation object

  - This will be used in about three places.
  - Events, Clubs, Organizations
  - We will have three different collections to handle these different events.
  - This is done for better scaling and separation of concerns
*/
type Invitation struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Type      string             `json:"type" bson:"type"`
	Sender    string             `json:"sender" bson:"sender"`
	Recipient string             `json:"recipient" bson:"recipient"`
	SubjectID primitive.ObjectID `json:"subject_id" bson:"subject_id"`
	Status    string             `json:"status" bson:"status"`
	Data      *InvitationData    `json:"data,omitempty" bson:"data,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

/*
Invitation Data

  - This holds data relevant to the invitation object
  - Based on the type only one of these variables will be populated at a time
*/
type InvitationData struct {
	Club         *Club         `json:"club,omitempty" bson:"club"`
	Event        *Event        `json:"event,omitempty" bson:"event"`
	Organization *Organization `json:"organization,omitempty" bson:"organization"`
}

type InvitationsResponse struct {
	TotalInvitations int          `json:"total_invitations"`
	Invitations      []Invitation `json:"invitations"`
}

type CheckIn struct {
	User          UserData        `json:"user" bson:"user"`
	Clubs         *[]Club         `json:"clubs,omitempty" bson:"clubs,omitempty"`
	Organizations *[]Organization `json:"organizations,omitempty" bson:"organizations,omitempty"`
	Invitations   *[]Invitation   `json:"invitations,omitempty" bson:"invitations,omitempty"`
	Token         *string         `json:"token,omitempty"`
}

type LocationResponse struct {
	Venues *[]Venue `json:"venues"`
	Events *[]Event `json:"events"`
}

type MemberDao struct {
	ID             primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         string              `json:"user_id" bson:"user_id"`
	Role           string              `json:"role" bson:"role"`
	ClubID         *primitive.ObjectID `json:"club_id,omitempty" bson:"club_id,omitempty"`
	OrganizationID *primitive.ObjectID `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	JoinedAt       primitive.DateTime  `json:"joined_at,omitempty" bson:"joined_at,omitempty"`
}
type Member struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	User     *UserSnippet       `json:"user,omitempty" bson:"user"`
	Role     string             `json:"role" bson:"role"`
	JoinedAt primitive.DateTime `json:"joined_at" bson:"joined_at"`
}

type Message struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Sender    *string             `json:"sender" bson:"sender"`
	Type      *string             `json:"type" bson:"type"`
	Body      *string             `json:"body" bson:"body"`
	Timestamp *primitive.DateTime `json:"timestamp" bson:"timestamp"`
}

type GroupVerificationDao struct {
	ID         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type       *string             `json:"type,omitempty" bson:"type,omitempty"`
	GroupID    *primitive.ObjectID `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Documents  *[]string           `json:"documents,omitempty" bson:"documents,omitempty"`
	VerifiedAt *primitive.DateTime `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
}

type UserVerificationDao struct {
	ID         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UUID       *string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Documents  *[]string           `json:"documents,omitempty" bson:"documents,omitempty"`
	VerifiedAt *primitive.DateTime `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
}

// Object to handle status update requests
type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

// Location data such as locale & coordinates
type Location struct {
	Country      string  `json:"country" bson:"country"`
	AdminArea    string  `json:"admin_area" bson:"admin_area"`
	SubAdminArea string  `json:"sub_admin_area" bson:"sub_admin_area"`
	Location     GeoJSON `json:"location" bson:"location"`
}

type GeoJSON struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type DeviceToken struct {
	Token  string `json:"token" bson:"token"`
	Device string `json:"device" bson:"device"`
}
