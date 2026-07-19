package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Invite type
type InviteType string

const (
	InviteTypeEvent InviteType = "EVENT"
	InviteTypeTeam  InviteType = "TEAM"
	InviteTypeClub  InviteType = "CLUB"
	InviteTypeOrg   InviteType = "ORG"
)

// Invite status
type InviteStatus string

const (
	InviteStatusPending  InviteStatus = "PENDING"
	InviteStatusAccepted InviteStatus = "ACCEPTED"
	InviteStatusDeclined InviteStatus = "DECLINED"
)

// Invite DB Object
type InviteRequest struct {
	ID          bson.ObjectID `bson:"_id"                json:"id"`
	Type        InviteType    `bson:"type"               json:"type"`
	ContextID   string        `bson:"context_id"         json:"context_id"`
	InviteeID   string        `bson:"invitee_id"         json:"invitee_id"`
	RequestorID string        `bson:"requestor_id"       json:"requestor_id"`
	Status      InviteStatus  `bson:"status"             json:"status"`
	UpdatedAt   bson.DateTime `bson:"updated_at"         json:"updated_at"`
	CreatedAt   bson.DateTime `bson:"created_at"         json:"created_at"`
}

// RabbitMQ message object
type InviteRequestMessage struct {
	ID          string     `json:"id"`
	Type        InviteType `json:"type"`
	ContextID   string     `json:"context_id"`
	RequestorID string     `json:"requestor_id"`
	InviteeIDs  []string   `json:"invitee_ids"`
	CreatedAt   time.Time  `json:"created_at"`
}

type InviteResponse struct {
	ID          string       `json:"id"`
	Type        InviteType   `json:"type"`
	ContextID   string       `json:"context_id"`
	InviteeID   string       `json:"invitee_id"`
	RequestorID string       `json:"requestor_id"`
	Status      InviteStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Single invite creation
type CreateInviteRequest struct {
	Type        InviteType `json:"type"`
	ContextID   string     `json:"context_id"`
	InviteeID   string     `json:"invitee_id"`
	RequestorID string     `json:"requestor_id"`
}

// Invite update request object
type UpdateInviteRequest struct {
	Status InviteStatus `json:"status"`
}

// User Invites response object
type UserInvitesResponse struct {
	Invites    []InviteResponse `json:"invites"`
	NextCursor string           `json:"next_cursor"`
}
