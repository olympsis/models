package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

/*
RSVPStatus was originally an integer enum (0/1/2). It is now a string enum, but
both wire formats are still in play, so this file gives the type explicit codecs
instead of relying on Go's default string encoding.

The rules, in one place:

  - MongoDB  — always written as the STRING ("YES"). Legacy integer documents
    still decode correctly, so a read never fails while the migration is in
    flight (see tools/migrate-rsvp-status.js).
  - REST API — always written as the LEGACY INTEGER (1). Shipped iOS/Android
    builds decode `status` into an Int-backed enum, so emitting a string there
    would break the participant list on every client already in the field.
    Requests are accepted in EITHER form, so new clients can start sending
    strings before the responses flip.
  - Event bus — always the STRING. It is a new, internal, server-to-service
    contract with no legacy consumers; see RSVPCreatedMessage.MarshalJSON.

When every shipped client understands strings, delete MarshalJSON below and the
REST responses become strings too — nothing else has to change.
*/

// Legacy integer codes for RSVPStatus, as stored/sent before the string switch.
// Only 0/1/2 ever existed on the wire; CAN'T is newer than the integer form and
// is assigned 3 so it round-trips rather than being collapsed into another
// status. Old clients will not recognize 3 — they also cannot produce it.
const (
	legacyRSVPMaybe    = 0
	legacyRSVPYes      = 1
	legacyRSVPWaitlist = 2
	legacyRSVPCant     = 3
)

// legacyRSVPCodes maps each status to its integer form and back.
var (
	legacyRSVPCodes = map[RSVPStatus]int{
		RSVPMaybe:    legacyRSVPMaybe,
		RSVPYes:      legacyRSVPYes,
		RSVPWaitlist: legacyRSVPWaitlist,
		RSVPCant:     legacyRSVPCant,
	}
	legacyRSVPStatuses = map[int]RSVPStatus{
		legacyRSVPMaybe:    RSVPMaybe,
		legacyRSVPYes:      RSVPYes,
		legacyRSVPWaitlist: RSVPWaitlist,
		legacyRSVPCant:     RSVPCant,
	}
)

// IsValid reports whether s is one of the four known statuses.
func (s RSVPStatus) IsValid() bool {
	_, ok := legacyRSVPCodes[s]
	return ok
}

// LegacyCode returns the integer form of the status. Unknown values fall back to
// MAYBE's code so a malformed record can't produce an out-of-range enum on a
// client.
func (s RSVPStatus) LegacyCode() int {
	if code, ok := legacyRSVPCodes[s]; ok {
		return code
	}
	return legacyRSVPMaybe
}

// RSVPStatusFromLegacyCode converts an integer status back to its string form.
func RSVPStatusFromLegacyCode(code int) (RSVPStatus, bool) {
	s, ok := legacyRSVPStatuses[code]
	return s, ok
}

// ParseRSVPStatus accepts either wire form as text: "YES", "yes", or "1".
// Used by query-parameter parsing, where everything arrives as a string.
func ParseRSVPStatus(v string) (RSVPStatus, error) {
	// Try the string form first, case-insensitively.
	s := RSVPStatus(strings.ToUpper(strings.TrimSpace(v)))
	if s.IsValid() {
		return s, nil
	}

	// Then the legacy integer form, e.g. "1". strconv.Atoi (not fmt.Sscanf) so
	// trailing garbage is rejected — Sscanf would happily read "2junk" as 2.
	if code, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
		if s, ok := RSVPStatusFromLegacyCode(code); ok {
			return s, nil
		}
	}
	return "", fmt.Errorf("invalid rsvp status: %q", v)
}

/*****
 JSON
*****/

// MarshalJSON emits the LEGACY INTEGER form so already-shipped clients keep
// decoding `status`. Remove this method (and only this method) to flip REST
// responses over to strings.
func (s RSVPStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.LegacyCode())
}

// UnmarshalJSON accepts both wire forms: `"YES"` (new clients) and `1` (clients
// already in the field). Anything else is an error rather than a silent default,
// so a typo surfaces as a 400 instead of an unintended RSVP.
func (s *RSVPStatus) UnmarshalJSON(data []byte) error {
	// String form.
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		parsed, err := ParseRSVPStatus(str)
		if err != nil {
			return err
		}
		*s = parsed
		return nil
	}

	// Legacy integer form.
	var code int
	if err := json.Unmarshal(data, &code); err == nil {
		parsed, ok := RSVPStatusFromLegacyCode(code)
		if !ok {
			return fmt.Errorf("invalid rsvp status code: %d", code)
		}
		*s = parsed
		return nil
	}

	return fmt.Errorf("invalid rsvp status: %s", string(data))
}

/*****
 BSON
*****/

// MarshalBSONValue always writes the STRING form, so every new/updated document
// lands in the target format regardless of what the client sent. This also
// covers query filters — bson.M{"status": RSVPWaitlist} serializes to "WAITLIST".
func (s RSVPStatus) MarshalBSONValue() (byte, []byte, error) {
	typ, data, err := bson.MarshalValue(string(s))
	return byte(typ), data, err
}

// UnmarshalBSONValue reads either form, so documents not yet converted by
// tools/migrate-rsvp-status.js still decode. Note that this only rescues READS —
// a query filtering on "WAITLIST" will not match a document still holding 2,
// which is why the migration is required rather than optional.
func (s *RSVPStatus) UnmarshalBSONValue(typ byte, data []byte) error {
	rv := bson.RawValue{Type: bson.Type(typ), Value: data}

	switch bson.Type(typ) {
	case bson.TypeString:
		str, ok := rv.StringValueOK()
		if !ok {
			return fmt.Errorf("rsvp status: malformed string value")
		}
		parsed, err := ParseRSVPStatus(str)
		if err != nil {
			return err
		}
		*s = parsed
		return nil

	case bson.TypeInt32, bson.TypeInt64, bson.TypeDouble:
		var code int
		if v, ok := rv.Int32OK(); ok {
			code = int(v)
		} else if v, ok := rv.Int64OK(); ok {
			code = int(v)
		} else if v, ok := rv.DoubleOK(); ok {
			code = int(v)
		} else {
			return fmt.Errorf("rsvp status: malformed numeric value")
		}
		parsed, ok := RSVPStatusFromLegacyCode(code)
		if !ok {
			return fmt.Errorf("invalid rsvp status code: %d", code)
		}
		*s = parsed
		return nil

	case bson.TypeNull, bson.TypeUndefined:
		// Leave the zero value; a nil/absent status is not an error.
		return nil
	}

	return fmt.Errorf("rsvp status: cannot decode from bson type %v", bson.Type(typ))
}
