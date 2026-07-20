package models

import (
	"encoding/json"
	"time"
)

// RSVPCreatedMessage is the JSON body published on the `rsvp.created` routing
// key by the main server and consumed by notif-service.
type RSVPCreatedMessage struct {
	ID        string     `json:"id"`         // participant record id (routes the tap)
	UserID    string     `json:"user_id"`    // the user who RSVPed
	EventID   string     `json:"event_id"`   // hex id of the event
	Status    RSVPStatus `json:"status"`     // RSVP status, always the string form on the bus
	CreatedAt time.Time  `json:"created_at"` // event time
}

// MarshalJSON forces `status` to its STRING form on the bus.
//
// RSVPStatus.MarshalJSON emits the legacy integer for REST responses (see
// rsvp.go), but the event bus is a new internal contract whose consumers expect
// "YES"/"MAYBE"/"CAN'T"/"WAITLIST". Overriding here keeps the two audiences from
// having to agree.
//
// The `alias` indirection strips this method from the embedded value, which is
// what stops MarshalJSON from calling itself; the outer Status field then
// shadows the embedded one because it sits at a shallower depth.
func (m RSVPCreatedMessage) MarshalJSON() ([]byte, error) {
	type alias RSVPCreatedMessage
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{
		alias:  alias(m),
		Status: string(m.Status),
	})
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
