package models

import "time"

// RSVPCreatedMessage is the JSON body this service expects on the `rsvp.created`
// routing key.
//
// IMPORTANT: the main server does NOT publish this today — it delivers RSVP
// pushes in-process. The server team must add a publisher that emits exactly
// these fields for the "New Participant" notification to fire.
type RSVPCreatedMessage struct {
	ID        string     `json:"id"`         // participant record id (routes the tap)
	UserID    string     `json:"user_id"`    // the user who RSVPed
	EventID   string     `json:"event_id"`   // hex id of the event
	Status    RSVPStatus `json:"status"`     // RSVP status
	CreatedAt time.Time  `json:"created_at"` // event time
}

// CommentCreatedMessage is the JSON body this service expects on the
// `comment.created` routing key. Like rsvp.created, the server does not publish
// this yet and must add a matching publisher.
type CommentCreatedMessage struct {
	ID        string    `json:"id"`         // comment id (routes the tap)
	UserID    string    `json:"user_id"`    // comment author
	EventID   string    `json:"event_id"`   // hex id of the event
	Text      string    `json:"text"`       // comment body (not used for the push; localized on-device)
	CreatedAt time.Time `json:"created_at"` // event time
}
