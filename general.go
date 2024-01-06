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
	CreatedAt int64              `json:"created_at" bson:"created_at"`
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
	User          UserData       `json:"user"`
	Clubs         []Club         `json:"clubs"`
	Organizations []Organization `json:"organizations"`
	Invitations   []Invitation   `json:"invitations"`
	Token         *string        `json:"token,omitempty"`
}
