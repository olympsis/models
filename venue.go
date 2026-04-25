package models

import "go.mongodb.org/mongo-driver/v2/bson"

/*
Venue

  - A physical location where sports activities happen.
  - A venue can own multiple bookable VenueUnits (courts, fields).
  - Availability here describes the venue's overall open hours; each unit
    can override with its own Availability.
  - units / transit_lines hold IDs into their own collections rather than
    embedding the full documents, so those can be paginated/queried on
    their own.
*/
type Venue struct {
	ID           bson.ObjectID `json:"id,omitempty" bson:"_id"`
	OwnerID      bson.ObjectID `json:"owner_id" bson:"owner_id"` // must be an organization
	Name         string        `json:"name" bson:"name"`
	Description  string        `json:"description" bson:"description"`
	Availability Availability  `json:"availability" bson:"availability"`

	Sports []string `json:"sports" bson:"sports"`
	Media  []string `json:"media" bson:"media"`
	URL    string   `json:"url" bson:"url"`

	// Address / location — locality & sub_locality are optional because
	// not every country / region model has them populated.
	Address            string  `json:"address" bson:"address"`
	Location           GeoJSON `json:"location" bson:"location"`
	Locality           *string `json:"locality,omitempty" bson:"locality,omitempty"`
	SubLocality        *string `json:"sub_locality,omitempty" bson:"sub_locality,omitempty"`
	AdministrativeArea string  `json:"administrative_area" bson:"administrative_area"`
	CountryCode        string  `json:"country_code" bson:"country_code"`
	Timezone           string  `json:"timezone" bson:"timezone"`

	// References to separate collections.
	Units        []bson.ObjectID `json:"units" bson:"units"`                 // bookable courts / fields
	TransitLines []bson.ObjectID `json:"transit_lines" bson:"transit_lines"` // nearby subway / bus lines

	Features     VenueFeatures     `json:"features" bson:"features"`
	Access       VenueAccess       `json:"access" bson:"access"`
	Capabilities VenueCapabilities `json:"capabilities" bson:"capabilities"`

	// Open-ended tags: "parking", "restrooms", "lockers", etc.
	Amenities []string `json:"amenities" bson:"amenities"`

	CreatedAt bson.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt *bson.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Physical attributes of the venue itself.
type VenueFeatures struct {
	Indoor      bool `json:"indoor" bson:"indoor"`
	Accessible  bool `json:"accessible" bson:"accessible"`
	Illuminated bool `json:"illuminated" bson:"illuminated"`
}

// Rules about who can use the venue / how.
type VenueAccess struct {
	RequiresPermit     bool `json:"requires_permit" bson:"requires_permit"`
	RequiresBooking    bool `json:"requires_booking" bson:"requires_booking"`
	RequiresMembership bool `json:"requires_membership" bson:"requires_membership"`
}

// App-side behavior the venue supports (queueing at public courts, etc.).
type VenueCapabilities struct {
	SupportsQueue     bool `json:"supports_queue" bson:"supports_queue"`
	SupportsWaitTimes bool `json:"supports_wait_times" bson:"supports_wait_times"`
}

type VenuesResponse struct {
	TotalVenues int     `json:"total_venues"`
	Venues      []Venue `json:"venues"`
}

/*
VenueDescriptor

  - A lightweight snapshot of a venue's identity + location.
  - Embedded on other documents (e.g. Event.Venues) so we don't have to
    re-fetch the full venue just to render a card.
*/
type VenueDescriptor struct {
	VenueID            bson.ObjectID `json:"venue_id" bson:"venue_id"`
	Name               string        `json:"name" bson:"name"`
	Location           GeoJSON       `json:"location" bson:"location"`
	Address            string        `json:"address" bson:"address"`
	Locality           string        `json:"locality" bson:"locality"`
	SubLocality        string        `json:"sub_locality" bson:"sub_locality"`
	AdministrativeArea string        `json:"administrative_area" bson:"administrative_area"`
	CountryCode        string        `json:"country_code" bson:"country_code"`
}

/*
Availability

  - Embedded on both Venue and VenueUnit.
  - RegularHours: weekly recurring open windows.
  - BlackoutDates: full-day closures (maintenance, holidays).
  - SpecialDates: one-off overrides with different open/close times.
*/
type Availability struct {
	RegularHours  []TimeSlot         `json:"regular_hours" bson:"regular_hours"`
	BlackoutDates []BlackoutTimeSlot `json:"blackout_dates" bson:"blackout_dates"`
	SpecialDates  []SpecialTimeSlot  `json:"special_dates" bson:"special_dates"`
}

// Day is a lowercase day name ("monday", "tuesday", ...).
// Open / Close are "HH:MM" in 24-hour format.
type TimeSlot struct {
	Day   string `json:"day" bson:"day"`
	Open  string `json:"open" bson:"open"`
	Close string `json:"close" bson:"close"`
}

// A full-day closure on a specific date.
type BlackoutTimeSlot struct {
	Date   bson.DateTime `json:"date" bson:"date"`
	Reason string        `json:"reason,omitempty" bson:"reason,omitempty"`
}

// An override of regular hours on a specific date (e.g. holiday hours).
type SpecialTimeSlot struct {
	Date   bson.DateTime `json:"date" bson:"date"`
	Open   string        `json:"open" bson:"open"`
	Close  string        `json:"close" bson:"close"`
	Reason string        `json:"reason,omitempty" bson:"reason,omitempty"`
}

/*
VenueUnit

  - A bookable sub-resource of a venue: a specific court, field, lane, etc.
  - Each unit has its own surface, rate schedule, and availability so a
    venue with mixed surfaces (e.g. 2 hard + 2 clay courts) is modeled
    naturally.
*/
type VenueUnit struct {
	ID           bson.ObjectID   `json:"id,omitempty" bson:"_id"`
	VenueID      bson.ObjectID   `json:"venue_id" bson:"venue_id"`
	Name         string          `json:"name" bson:"name"`
	UnitType     string          `json:"unit_type" bson:"unit_type"` // "court", "field", etc.
	Location     GeoJSON         `json:"location" bson:"location"`
	Surface      Surface         `json:"surface" bson:"surface"`
	Sports       []string        `json:"sports" bson:"sports"`
	Rates        []VenueUnitRate `json:"rates" bson:"rates"`
	Availability Availability    `json:"availability" bson:"availability"`
}

/*
VenueUnitRate

  - Price for a unit for a day + time-range combination.
  - RateMinor is in the smallest currency unit (cents for USD) to avoid
    floating-point rounding.
  - TimeRange is a string like "09:00-17:00" — parsing lives in app code.
*/
type VenueUnitRate struct {
	Day       string `json:"day" bson:"day"`
	TimeRange string `json:"time_range" bson:"time_range"`
	Currency  string `json:"currency" bson:"currency"`
	RateMinor int64  `json:"rate_minor" bson:"rate_minor"`
}

/*
TransitLine

  - A public-transit line (subway, bus) serving the venue.
  - Stored in its own collection so the same line (e.g. the NYC "Q")
    can be referenced by many venues without duplication.
*/
type TransitLine struct {
	ID                 bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Type               string        `json:"type" bson:"type"` // "subway", "bus", "train", ...
	Name               string        `json:"name" bson:"name"`
	System             string        `json:"system" bson:"system"` // e.g. "MTA"
	Color              string        `json:"color" bson:"color"`
	IconURL            string        `json:"icon_url" bson:"icon_url"`
	Locality           string        `json:"locality" bson:"locality"`
	AdministrativeArea string        `json:"administrative_area" bson:"administrative_area"`
	CountryCode        string        `json:"country_code" bson:"country_code"`
}

/*
VenueReservation

  - A booking of a specific VenueUnit for a time range.
  - Payment flows through Stripe — TransactionID is the Stripe payment id.
  - A reservation is created in "pending" state with ExpiresAt set; a TTL
    index on expires_at cleans up abandoned bookings. On successful
    payment the webhook flips status to "confirmed" and clears ExpiresAt.
  - ClubID / OrganizationID are optional — set when the booking is on
    behalf of a club / org rather than an individual.
*/
type VenueReservation struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id"`
	VenueID     bson.ObjectID `json:"venue_id" bson:"venue_id"`
	VenueUnitID bson.ObjectID `json:"venue_unit_id" bson:"venue_unit_id"`

	// Stripe payment intent / charge id. Unique index lives on this field
	// so webhook replays are idempotent.
	TransactionID string `json:"transaction_id" bson:"transaction_id"`

	UserID         string         `json:"user_id" bson:"user_id"`
	ClubID         *bson.ObjectID `json:"club_id,omitempty" bson:"club_id,omitempty"`
	OrganizationID *bson.ObjectID `json:"organization_id,omitempty" bson:"organization_id,omitempty"`

	StartDate bson.DateTime `json:"start_date" bson:"start_date"`
	EndDate   bson.DateTime `json:"end_date" bson:"end_date"`
	Timezone  string        `json:"timezone" bson:"timezone"`

	Currency        string `json:"currency" bson:"currency"`
	AmountPaidMinor int64  `json:"amount_paid_minor" bson:"amount_paid_minor"`

	Status ReservationStatus `json:"status" bson:"status"`
	// Only set while Status == pending — drives the TTL cleanup.
	ExpiresAt *bson.DateTime `json:"expires_at,omitempty" bson:"expires_at,omitempty"`

	CreatedBy string         `json:"created_by" bson:"created_by"`
	CreatedAt bson.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt *bson.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type VenueCreationRequest struct {
	Venue      Venue       `json:"venue"`
	VenueUnits []VenueUnit `json:"venue_units"`
}
