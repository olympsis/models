package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Poster     UserSnippet        `json:"poster" bson:"poster"`
	Organizers []Organizer        `json:"organizers" bson:"organizers"`
	Venues     []VenueDescriptor  `json:"venues" bson:"venues"`

	MediaURL  string    `json:"media_url" bson:"media_url"`
	MediaType MediaType `json:"media_type" bson:"media_type"`

	Title  string   `json:"title" bson:"title"`
	Body   string   `json:"body" bson:"body"`
	Sports []string `json:"sports" bson:"sports"`
	Tags   []string `json:"tags" bson:"tags"`

	FormatConfig *EventFormatConfig `json:"format_config,omitempty" bson:"format_config,omitempty"`

	StartTime primitive.DateTime `json:"start_time" bson:"start_time"`
	StopTime  primitive.DateTime `json:"stop_time" bson:"stop_time"`

	Participants       []Participant       `json:"participants,omitempty" bson:"participants,omitempty"`
	ParticipantsConfig *ParticipantsConfig `json:"participants_config" bson:"participants_config"`

	Teams       []Team       `json:"teams,omitempty" bson:"teams,omitempty"`
	TeamsConfig *TeamsConfig `json:"teams_config,omitempty" bson:"teams_config,omitempty"`

	Visibility   VisibilityScope `json:"visibility" bson:"visibility"`
	ExternalLink *string         `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  bool            `json:"is_sensitive" bson:"is_sensitive"`

	CreatedAt   primitive.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime  `json:"updated_at" bson:"updated_at"`
	CancelledAt *primitive.DateTime `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`

	RecurrenceConfig *EventRecurrenceConfig `json:"recurrence_config,omitempty" bson:"recurrence_config,omitempty"`
}

type EventDao struct {
	PosterID   *string            `json:"poster_id,omitempty" bson:"poster_id,omitempty"`
	Organizers *[]Organizer       `json:"organizers,omitempty" bson:"organizers,omitempty"`
	Venues     *[]VenueDescriptor `json:"venues,omitempty" bson:"venues,omitempty"`

	MediaURL  string    `json:"media_url,omitempty" bson:"media_url,omitempty"`
	MediaType MediaType `json:"media_type,omitempty" bson:"media_type,omitempty"`

	Title  *string   `json:"title,omitempty" bson:"title,omitempty"`
	Body   *string   `json:"body,omitempty" bson:"body,omitempty"`
	Sports *[]string `json:"sports,omitempty" bson:"sports,omitempty"`
	Tags   *[]string `json:"tags,omitempty" bson:"tags,omitempty"`

	FormatConfig *EventFormatConfig `json:"format_config,omitempty" bson:"format_config,omitempty"`

	StartTime *primitive.DateTime `json:"start_time,omitempty" bson:"start_time,omitempty"`
	StopTime  *primitive.DateTime `json:"stop_time,omitempty" bson:"stop_time,omitempty"`

	ParticipantsConfig *ParticipantsConfig `json:"participants_config" bson:"participants_config"`

	TeamsConfig *TeamsConfig `json:"teams_config,omitempty" bson:"teams_config,omitempty"`

	Visibility   *VisibilityScope `json:"visibility,omitempty" bson:"visibility,omitempty"`
	ExternalLink *string          `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  *bool            `json:"is_sensitive,omitempty" bson:"is_sensitive,omitempty"`

	CreatedAt   *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CancelledAt *primitive.DateTime `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`

	RecurrenceConfig *EventRecurrenceConfig `json:"recurrence_config,omitempty" bson:"recurrence_config,omitempty"`
}

type EventFormatConfig struct {
	// Type identifiers
	IsCompetition       *bool               `json:"is_competition,omitempty" bson:"is_competition,omitempty"`
	IsCompetitionGame   *bool               `json:"is_competition_game,omitempty" bson:"is_competition_game,omitempty"`
	ParentCompetitionID *primitive.ObjectID `json:"parent_competition_id,omitempty" bson:"parent_competition_id,omitempty"`
	CompetitionState    *string             `json:"competition_state,omitempty" bson:"competition_state,omitempty"` // "not_started", "in_progress", "completed"

	// Format details
	Format       *CompetitionFormats `json:"format,omitempty" bson:"format,omitempty"`
	Rounds       *int32              `json:"rounds,omitempty" bson:"rounds,omitempty"`
	CurrentRound *int32              `json:"current_round,omitempty" bson:"current_round,omitempty"`
	BracketData  *any                `json:"bracket_data,omitempty" bson:"bracket_data,omitempty"`

	// Registration period (if you decide to include it here)
	RegistrationStart     *primitive.DateTime `json:"registration_start,omitempty" bson:"registration_start,omitempty"`
	RegistrationEnd       *primitive.DateTime `json:"registration_end,omitempty" bson:"registration_end,omitempty"`
	AllowLateRegistration *bool               `json:"allow_late_registration,omitempty" bson:"allow_late_registration,omitempty"`
}

type EventRecurrenceConfig struct {
	RecurrenceRule   *string               `json:"recurrence_rule,omitempty" bson:"recurrence_rule,omitempty"`
	RecurrenceEnd    *primitive.DateTime   `json:"recurrence_end,omitempty" bson:"recurrence_end,omitempty"`
	ParentEventID    *primitive.ObjectID   `json:"parent_event_id,omitempty" bson:"parent_event_id,omitempty"`
	DeletedInstances *[]primitive.ObjectID `json:"deleted_instances,omitempty" bson:"deleted_instances,omitempty"`
}

type EventsResponse struct {
	TotalEvents int16   `json:"total_events"`
	Events      []Event `json:"events"`
}

type EventData struct {
	Comments []Comment `json:"comments"`

	Teams         []Team `json:"teams"`
	TeamsWaitlist []Team `json:"teams_waitlist"`

	Participants []Participant `json:"participants"`
	Waitlist     []Participant `json:"waitlist"`
}

type Participant struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Status    RSVPStatus         `json:"status" bson:"status"`
	EventID   primitive.ObjectID `json:"event_id" bson:"event_id"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at"`
}

type ParticipantDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    *string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Status    *RSVPStatus         `json:"status,omitempty" bson:"status,omitempty"`
	EventID   *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type ParticipantsConfig struct {
	HasWaitlist     *bool  `json:"has_waitlist,omitempty" bson:"has_waitlist,omitempty"`
	MinParticipants *int32 `json:"min_participants,omitempty" bson:"min_participants,omitempty"`
	MaxParticipants *int32 `json:"max_participants,omitempty" bson:"max_participants,omitempty"`
}

type Team struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Members   []Participant      `json:"members" bson:"members"`
	EventID   primitive.ObjectID `json:"event_id" bson:"event_id"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at"`
}

type TeamDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      *string             `json:"name,omitempty" bson:"name,omitempty"`
	Members   *[]Participant      `json:"members,omitempty" bson:"members,omitempty"`
	EventID   *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type TeamsConfig struct {
	HasWaitlist *bool  `json:"has_waitlist,omitempty" bson:"has_waitlist,omitempty"`
	MinTeams    *int32 `json:"min_teams,omitempty" bson:"min_teams,omitempty"`
	MaxTeams    *int32 `json:"max_teams,omitempty" bson:"max_teams,omitempty"`
	MaxTeamSize *int32 `json:"max_team_size,omitempty" bson:"max_team_size,omitempty"`
}

type FullEventError struct {
	MSG   string `json:"msg"`
	Event Event  `json:"event"`
}

type Organizer struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type int8               `json:"type" bson:"type"`
}

type NewEventDao struct {
	Event       EventDao           `json:"event"`
	IncludeHost *bool              `json:"include_host,omitempty"`
	Recurrence  *RecurrenceOptions `json:"recurrence,omitempty"`
}

type RecurrenceOptions struct {
	Pattern  string             `json:"pattern"`  // "DAILY", "WEEKLY", "MONTHLY"
	EndTime  primitive.DateTime `json:"end_time"` // Must be after the event's start_time
	Interval int                `json:"interval"` // e.g., every 2 weeks
}

type EventComment struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Text      string             `json:"text" bson:"text"`
	EventID   primitive.ObjectID `json:"event_id" bson:"event_id"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

type EventCommentDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UserID    *string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Text      *string             `json:"text,omitempty" bson:"text,omitempty"`
	EventID   primitive.ObjectID  `json:"event_id,omitempty" bson:"event_id,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// 						//
// EVENT LOGGING MODELS //
//						//

type EventAuditLog struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	EventID       primitive.ObjectID `json:"event_id" bson:"event_id"`
	UserID        string             `json:"user_id" bson:"user_id"`
	Action        string             `json:"action" bson:"action"` // "create", "update", "cancel", etc.
	FieldsChanged []string           `json:"fields_changed,omitempty" bson:"fields_changed,omitempty"`
	OldValues     map[string]any     `json:"old_values,omitempty" bson:"old_values,omitempty"`
	NewValues     map[string]any     `json:"new_values,omitempty" bson:"new_values,omitempty"`
	Timestamp     primitive.DateTime `json:"timestamp" bson:"timestamp"`
}

type EventAnalytics struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	EventID primitive.ObjectID `json:"event_id" bson:"event_id"`

	// Overview metrics
	ViewCount         int32 `json:"view_count" bson:"view_count"`
	UniqueViewerCount int32 `json:"unique_viewer_count" bson:"unique_viewer_count"`

	// Conversion metrics
	RSVPCount      int32   `json:"rsvp_count" bson:"rsvp_count"`
	WaitlistCount  int32   `json:"waitlist_count" bson:"waitlist_count"`
	ConversionRate float32 `json:"conversion_rate" bson:"conversion_rate"` // ViewToRSVP %

	// Engagement metrics
	ShareCount   int32 `json:"share_count" bson:"share_count"`
	SaveCount    int32 `json:"save_count" bson:"save_count"`
	CommentCount int32 `json:"comment_count" bson:"comment_count"`

	// Time-series data (for charts)
	DailyMetrics map[string]*DailyEngagement `json:"daily_metrics" bson:"daily_metrics"` // Date string -> metrics

	// Audience demographics (for pie charts)
	PlatformBreakdown map[string]int32 `json:"platform_breakdown" bson:"platform_breakdown"` // "ios", "android", "web"
	SourceBreakdown   map[string]int32 `json:"source_breakdown" bson:"source_breakdown"`     // "search", "feed", "direct"

	// For competition events
	TeamViewMetrics map[string]int32 `json:"team_view_metrics,omitempty" bson:"team_view_metrics,omitempty"` // TeamID -> views

	// Update tracking
	LastUpdated primitive.DateTime `json:"last_updated" bson:"last_updated"`
}

type DailyEngagement struct {
	Views         int32 `json:"views" bson:"views"`
	UniqueViews   int32 `json:"unique_views" bson:"unique_views"`
	RSVPs         int32 `json:"rsvps" bson:"rsvps"`
	Shares        int32 `json:"shares" bson:"shares"`
	AvgTimeOnPage int32 `json:"avg_time_on_page" bson:"avg_time_on_page"` // seconds
}

type EventDetailedView struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	EventID primitive.ObjectID `json:"event_id" bson:"event_id"`

	// User info (partially anonymized)
	UserID   *string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserType string  `json:"user_type" bson:"user_type"` // "member", "non-member", "anonymous"

	// Session details
	ViewTime primitive.DateTime  `json:"view_time" bson:"view_time"`
	ExitTime *primitive.DateTime `json:"exit_time,omitempty" bson:"exit_time,omitempty"`
	Duration int32               `json:"duration" bson:"duration"` // Time in seconds

	// View context
	Source    string `json:"source" bson:"source"`                           // "search", "profile", "feed", "notification"
	EntryPath string `json:"entry_path" bson:"entry_path"`                   // URL path they came from
	ExitPath  string `json:"exit_path,omitempty" bson:"exit_path,omitempty"` // Where they went next

	// Technical data
	DeviceType string `json:"device_type" bson:"device_type"` // "mobile", "tablet", "desktop"
	Platform   string `json:"platform" bson:"platform"`       // "ios", "android", "web"
	Country    string `json:"country,omitempty" bson:"country,omitempty"`
	Region     string `json:"region,omitempty" bson:"region,omitempty"`
}
