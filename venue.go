package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Venue struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Owner       Ownership          `json:"owner" bson:"owner"`
	Description string             `json:"description" bson:"description"`
	Sports      []string           `json:"sports" bson:"sports"`
	Images      []string           `json:"images" bson:"images"`

	// Location data
	City        string  `json:"city" bson:"city"`
	State       string  `json:"state" bson:"state"`
	Country     string  `json:"country" bson:"country"`
	Location    GeoJSON `json:"location" bson:"location"`
	SnapshotURL *string `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`

	// Booking info
	RequiresBooking bool    `json:"requires_booking" bson:"requires_booking"`
	BookingURL      *string `json:"booking_url,omitempty" bson:"booking_url,omitempty"`

	// Embedded capacity info (relatively static and small)
	Capacity *VenueCapacity `json:"capacity,omitempty" bson:"capacity,omitempty"`

	// Regular availability (weekly schedule - relatively static)
	WeeklySchedule *[]DayAvailability `json:"weekly_schedule,omitempty" bson:"weekly_schedule,omitempty"`

	// Reference IDs to separate collections (for data that changes frequently)
	BlackoutDatesID *primitive.ObjectID `json:"blackout_dates_id,omitempty" bson:"blackout_dates_id,omitempty"`
	SpecialHoursID  *primitive.ObjectID `json:"special_hours_id,omitempty" bson:"special_hours_id,omitempty"`

	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Ownership struct {
	Name    string              `json:"name" bson:"name"`
	Type    string              `json:"type" bson:"type"`
	OwnerID *primitive.ObjectID `json:"owner_id,omitempty" bson:"owner_id,omitempty"`
}

type VenuesResponse struct {
	TotalVenues int     `json:"total_venues"`
	Venues      []Venue `json:"venues"`
}

type VenueDescriptor struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     *string             `json:"name,omitempty" bson:"name,omitempty"`
	City     *string             `json:"city,omitempty" bson:"city,omitempty"`
	State    *string             `json:"state,omitempty" bson:"state,omitempty"`
	Country  *string             `json:"country,omitempty" bson:"country,omitempty"`
	Location *GeoJSON            `json:"location,omitempty" bson:"location,omitempty"`
}

type VenueAvailability struct {
	ID             primitive.ObjectID  `json:"id" bson:"_id"`
	VenueID        primitive.ObjectID  `json:"venue_id" bson:"venue_id"`
	WeeklySchedule []DayAvailability   `json:"weekly_schedule" bson:"weekly_schedule"`
	BlackoutDates  []BlackoutPeriod    `json:"blackout_dates" bson:"blackout_dates"`
	SpecialHours   []SpecialHourPeriod `json:"special_hours,omitempty" bson:"special_hours,omitempty"`
	UpdatedAt      primitive.DateTime  `json:"updated_at" bson:"updated_at"`
}

type DayAvailability struct {
	DayOfWeek   int             `json:"day_of_week" bson:"day_of_week"` // 0-6, 0 = Sunday
	IsAvailable bool            `json:"is_available" bson:"is_available"`
	TimeSlots   []TimeRangeSlot `json:"time_slots" bson:"time_slots"`
}

type TimeRangeSlot struct {
	StartTime string `json:"start_time" bson:"start_time"` // Format: "HH:MM" in 24h
	EndTime   string `json:"end_time" bson:"end_time"`     // Format: "HH:MM" in 24h
}

type BlackoutPeriod struct {
	StartDate      primitive.DateTime `json:"start_date" bson:"start_date"`
	EndDate        primitive.DateTime `json:"end_date" bson:"end_date"`
	Reason         string             `json:"reason,omitempty" bson:"reason,omitempty"`
	IsRecurring    bool               `json:"is_recurring" bson:"is_recurring"` // For annual events
	RecurrenceRule string             `json:"recurrence_rule,omitempty" bson:"recurrence_rule,omitempty"`
}

type SpecialHourPeriod struct {
	Date        primitive.DateTime `json:"date" bson:"date"`
	TimeSlots   []TimeRangeSlot    `json:"time_slots" bson:"time_slots"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}

type VenueCapacity struct {
	ID               primitive.ObjectID       `json:"id" bson:"_id"`
	VenueID          primitive.ObjectID       `json:"venue_id" bson:"venue_id"`
	TotalCapacity    int32                    `json:"total_capacity" bson:"total_capacity"`                             // Maximum number of people
	PerSportCapacity map[string]SportCapacity `json:"per_sport_capacity,omitempty" bson:"per_sport_capacity,omitempty"` // Sport-specific capacities
	Areas            []VenueArea              `json:"areas,omitempty" bson:"areas,omitempty"`                           // Optional breakdown by area
	UpdatedAt        primitive.DateTime       `json:"updated_at" bson:"updated_at"`
}

type SportCapacity struct {
	MaxParticipants int32  `json:"max_participants" bson:"max_participants"`
	MaxTeams        int32  `json:"max_teams,omitempty" bson:"max_teams,omitempty"`
	MaxSpectators   int32  `json:"max_spectators,omitempty" bson:"max_spectators,omitempty"`
	Notes           string `json:"notes,omitempty" bson:"notes,omitempty"`
}

type VenueArea struct {
	Name            string   `json:"name" bson:"name"` // e.g., "Court 1", "Main Field"
	Capacity        int32    `json:"capacity" bson:"capacity"`
	SupportedSports []string `json:"supported_sports,omitempty" bson:"supported_sports,omitempty"`
}

type VenueEventLog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	VenueID   primitive.ObjectID `json:"venue_id" bson:"venue_id"`
	EventID   primitive.ObjectID `json:"event_id" bson:"event_id"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}
