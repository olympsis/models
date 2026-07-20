package models

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// TestRSVPStatusJSONMarshalsLegacyInteger locks in the REST contract: responses
// carry the integer form so already-shipped clients keep decoding.
func TestRSVPStatusJSONMarshalsLegacyInteger(t *testing.T) {
	cases := map[RSVPStatus]string{
		RSVPMaybe:    "0",
		RSVPYes:      "1",
		RSVPWaitlist: "2",
		RSVPCant:     "3",
	}
	for status, want := range cases {
		got, err := json.Marshal(status)
		if err != nil {
			t.Fatalf("marshal %q: %v", status, err)
		}
		if string(got) != want {
			t.Errorf("marshal %q = %s, want %s", status, got, want)
		}
	}
}

// TestRSVPStatusJSONUnmarshalAcceptsBothForms is the backwards-compatibility
// guarantee for request bodies: old clients send 1, new clients send "YES".
func TestRSVPStatusJSONUnmarshalAcceptsBothForms(t *testing.T) {
	cases := map[string]RSVPStatus{
		`1`:          RSVPYes,
		`"YES"`:      RSVPYes,
		`"yes"`:      RSVPYes, // case-insensitive
		`0`:          RSVPMaybe,
		`2`:          RSVPWaitlist,
		`"CAN'T"`:    RSVPCant,
		`3`:          RSVPCant,
		`"MAYBE"`:    RSVPMaybe,
		`"WAITLIST"`: RSVPWaitlist,
	}
	for input, want := range cases {
		var got RSVPStatus
		if err := json.Unmarshal([]byte(input), &got); err != nil {
			t.Fatalf("unmarshal %s: %v", input, err)
		}
		if got != want {
			t.Errorf("unmarshal %s = %q, want %q", input, got, want)
		}
	}
}

// TestRSVPStatusJSONUnmarshalRejectsGarbage — an unrecognized status must be a
// hard error, not a silent default to MAYBE.
func TestRSVPStatusJSONUnmarshalRejectsGarbage(t *testing.T) {
	for _, input := range []string{`"NOPE"`, `99`, `-1`, `{}`, `true`} {
		var got RSVPStatus
		if err := json.Unmarshal([]byte(input), &got); err == nil {
			t.Errorf("unmarshal %s: expected error, got %q", input, got)
		}
	}
}

// TestRSVPStatusBSONWritesString is the whole point of the migration: whatever
// comes in, Mongo gets the string.
func TestRSVPStatusBSONWritesString(t *testing.T) {
	type doc struct {
		Status RSVPStatus `bson:"status"`
	}
	data, err := bson.Marshal(doc{Status: RSVPWaitlist})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw bson.Raw = data
	val, err := raw.LookupErr("status")
	if err != nil {
		t.Fatalf("lookup status: %v", err)
	}
	if val.Type != bson.TypeString {
		t.Fatalf("status stored as %v, want string", val.Type)
	}
	if s, _ := val.StringValueOK(); s != "WAITLIST" {
		t.Errorf("status = %q, want %q", s, "WAITLIST")
	}
}

// TestRSVPStatusBSONReadsLegacyInteger covers documents not yet converted by
// tools/migrate-rsvp-status.js — reads must not fail mid-migration.
func TestRSVPStatusBSONReadsLegacyInteger(t *testing.T) {
	type doc struct {
		Status RSVPStatus `bson:"status"`
	}

	for code, want := range map[int32]RSVPStatus{
		0: RSVPMaybe,
		1: RSVPYes,
		2: RSVPWaitlist,
	} {
		data, err := bson.Marshal(bson.M{"status": code})
		if err != nil {
			t.Fatalf("marshal legacy %d: %v", code, err)
		}
		var got doc
		if err := bson.Unmarshal(data, &got); err != nil {
			t.Fatalf("unmarshal legacy %d: %v", code, err)
		}
		if got.Status != want {
			t.Errorf("legacy %d decoded to %q, want %q", code, got.Status, want)
		}
	}
}

// TestRSVPStatusBSONReadsString — the post-migration happy path.
func TestRSVPStatusBSONReadsString(t *testing.T) {
	type doc struct {
		Status RSVPStatus `bson:"status"`
	}
	data, err := bson.Marshal(bson.M{"status": "CAN'T"})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got doc
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Status != RSVPCant {
		t.Errorf("decoded %q, want %q", got.Status, RSVPCant)
	}
}

// TestRSVPStatusPointerFieldRoundTrips — ParticipantDao.Status is a *RSVPStatus,
// so exercise that shape explicitly rather than only the value form.
func TestRSVPStatusPointerFieldRoundTrips(t *testing.T) {
	data, err := bson.Marshal(bson.M{"status": int32(2)})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got ParticipantDao
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Status == nil || *got.Status != RSVPWaitlist {
		t.Fatalf("status = %v, want %q", got.Status, RSVPWaitlist)
	}

	// And back out as a string.
	out, err := bson.Marshal(got)
	if err != nil {
		t.Fatalf("re-marshal: %v", err)
	}
	val, err := bson.Raw(out).LookupErr("status")
	if err != nil {
		t.Fatalf("lookup: %v", err)
	}
	if val.Type != bson.TypeString {
		t.Errorf("re-marshalled as %v, want string", val.Type)
	}
}

// TestRSVPCreatedMessageStatusIsString guards the bus contract, which differs
// from the REST contract on purpose.
func TestRSVPCreatedMessageStatusIsString(t *testing.T) {
	data, err := json.Marshal(RSVPCreatedMessage{
		ID:      "abc",
		UserID:  "user",
		EventID: "event",
		Status:  RSVPYes,
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded["status"] != "YES" {
		t.Errorf("bus status = %v, want \"YES\"", decoded["status"])
	}

	// And it must decode back into the typed field.
	var msg RSVPCreatedMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		t.Fatalf("re-unmarshal: %v", err)
	}
	if msg.Status != RSVPYes {
		t.Errorf("round-tripped status = %q, want %q", msg.Status, RSVPYes)
	}
}
