package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// TestTeamMemberRoleRoundTrips confirms the new Role field survives a BSON
// round-trip on both the value (Team) and pointer (TeamDao) shapes, and that it
// is stored as the plain string the MemberRole enum represents.
func TestTeamMemberRoleRoundTrips(t *testing.T) {
	uid := "user-1"
	role := OwnerMember
	dao := TeamMemberDao{UserID: &uid, Role: &role}

	data, err := bson.Marshal(dao)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	// Stored as a string.
	val, err := bson.Raw(data).LookupErr("role")
	if err != nil {
		t.Fatalf("lookup role: %v", err)
	}
	if val.Type != bson.TypeString {
		t.Fatalf("role stored as %v, want string", val.Type)
	}
	if s, _ := val.StringValueOK(); s != "OWNER" {
		t.Errorf("role = %q, want %q", s, "OWNER")
	}

	// Decodes back into the value shape used for reads.
	var got TeamMember
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Role != OwnerMember {
		t.Errorf("decoded role = %q, want %q", got.Role, OwnerMember)
	}
}

// TestTeamIsOpenRoundTrips guards the open/closed flag both when it is set and,
// on the Dao, when it is left nil (omitempty must drop it so a partial update
// doesn't clobber the stored value).
func TestTeamIsOpenRoundTrips(t *testing.T) {
	open := true
	name := "Sharks"
	dao := TeamDao{Name: &name, IsOpen: &open}

	data, err := bson.Marshal(dao)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got Team
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !got.IsOpen {
		t.Errorf("is_open = %v, want true", got.IsOpen)
	}

	// nil IsOpen must be omitted entirely.
	empty, err := bson.Marshal(TeamDao{Name: &name})
	if err != nil {
		t.Fatalf("marshal empty: %v", err)
	}
	if _, err := bson.Raw(empty).LookupErr("is_open"); err == nil {
		t.Error("is_open should be omitted when nil")
	}
}

// TestTeamsConfigRequiredRoundTrips covers the per-event RSVP-mode switch,
// including that nil is omitted (so it reads as individual-RSVP by default).
func TestTeamsConfigRequiredRoundTrips(t *testing.T) {
	required := true
	data, err := bson.Marshal(TeamsConfig{Required: &required})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got TeamsConfig
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Required == nil || !*got.Required {
		t.Errorf("required = %v, want true", got.Required)
	}

	empty, err := bson.Marshal(TeamsConfig{})
	if err != nil {
		t.Fatalf("marshal empty: %v", err)
	}
	if _, err := bson.Raw(empty).LookupErr("required"); err == nil {
		t.Error("required should be omitted when nil")
	}
}

// TestTeamApplicationStatusIsTypedString ensures TeamApplication uses the typed
// ApplicationStatus enum stored as its uppercase string form. This is the
// regression guard against repeating the club application's lowercase-vs-enum
// mismatch bug.
func TestTeamApplicationStatusIsTypedString(t *testing.T) {
	applicant := "user-2"
	status := PendingApplicationStatus
	dao := TeamApplicationDao{Applicant: &applicant, Status: &status}

	data, err := bson.Marshal(dao)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	val, err := bson.Raw(data).LookupErr("status")
	if err != nil {
		t.Fatalf("lookup status: %v", err)
	}
	if val.Type != bson.TypeString {
		t.Fatalf("status stored as %v, want string", val.Type)
	}
	if s, _ := val.StringValueOK(); s != "PENDING" {
		t.Errorf("status = %q, want %q", s, "PENDING")
	}

	// Round-trips through the stored (Dao) shape — the read shape's Applicant is
	// an aggregation-hydrated *UserSnippet, so a raw stored doc is only ever
	// decoded back into the Dao (its Applicant is the raw *string uid).
	var got TeamApplicationDao
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Status == nil || *got.Status != PendingApplicationStatus {
		t.Errorf("decoded status = %v, want %q", got.Status, PendingApplicationStatus)
	}
	if got.Applicant == nil || *got.Applicant != applicant {
		t.Errorf("decoded applicant = %v, want %q", got.Applicant, applicant)
	}
}
