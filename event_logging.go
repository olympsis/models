package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// EventAuditLog is a MongoDB time-series model.
// Handles the logging and tracing of event changes.
//
// Create collection with:
//
//	opts := options.CreateCollection().
//		SetTimeSeriesOptions(options.TimeSeries().
//			SetTimeField("timestamp").
//			SetMetaField("event_id").
//			SetGranularity("seconds")).
//		SetExpireAfterSeconds(7776000) // 90 days
//
//	err := db.CreateCollection(ctx, "event_audit_logs", opts)
type EventAuditLog struct {
	EventID primitive.ObjectID `json:"event_id" bson:"event_id"`

	// User & action tracing
	UserID string `json:"user_id" bson:"user_id"`
	Action string `json:"action" bson:"action"` // "create", "update", "cancel", etc.

	// Values changed
	FieldsChanged []string       `json:"fields_changed,omitempty" bson:"fields_changed,omitempty"`
	OldValues     map[string]any `json:"old_values,omitempty" bson:"old_values,omitempty"`
	NewValues     map[string]any `json:"new_values,omitempty" bson:"new_values,omitempty"`

	// Time of the event action
	Timestamp primitive.DateTime `json:"timestamp" bson:"timestamp"`
}

// EventViewLog is a MongoDB time-series model.
// Handles the logging and tracing of event view events.
// Documents are inserted on session exit with all fields populated.
//
// Create collection with:
//
//	opts := options.CreateCollection().
//		SetTimeSeriesOptions(options.TimeSeries().
//			SetTimeField("view_time").
//			SetMetaField("event_id").
//			SetGranularity("minutes")).
//		SetExpireAfterSeconds(7776000) // 90 days
//
//	err := db.CreateCollection(ctx, "event_view_logs", opts)
type EventViewLog struct {
	EventID primitive.ObjectID `json:"event_id" bson:"event_id"`

	// User info (partially anonymized)
	UserID   *string  `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserType UserType `json:"user_type" bson:"user_type"` // "member", "non-member", "anonymous"

	// Session details (calculated client-side before insert)
	ViewTime primitive.DateTime `json:"view_time" bson:"view_time"` // TimeField for time-series
	ExitTime primitive.DateTime `json:"exit_time" bson:"exit_time"`
	Duration int32              `json:"duration" bson:"duration"` // seconds

	// View context
	Source    *string `json:"source,omitempty" bson:"source,omitempty"`         // "search", "profile", "feed", "notification"
	EntryPath *string `json:"entry_path,omitempty" bson:"entry_path,omitempty"` // URL path they came from
	ExitPath  *string `json:"exit_path,omitempty" bson:"exit_path,omitempty"`   // Where they went next

	// Technical data
	DeviceType string  `json:"device_type" bson:"device_type"` // "mobile", "tablet", "desktop"
	Platform   string  `json:"platform" bson:"platform"`       // "ios", "android", "web"
	Country    *string `json:"country,omitempty" bson:"country,omitempty"`
	Region     *string `json:"region,omitempty" bson:"region,omitempty"`
}

// EventAnalytics is a computed response model for event metrics.
// Aggregated on fetch from EventViewLog and EventAuditLog time-series data.
// Not persisted - generated via MongoDB aggregation pipelines.
type EventAnalytics struct {
	EventID primitive.ObjectID `json:"event_id" bson:"event_id"`

	// Overview metrics
	ViewCount         int32 `json:"view_count" bson:"view_count"`
	UniqueViewerCount int32 `json:"unique_viewer_count" bson:"unique_viewer_count"`

	// Conversion metrics
	RSVPCount     int32 `json:"rsvp_count" bson:"rsvp_count"`
	WaitlistCount int32 `json:"waitlist_count" bson:"waitlist_count"`

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

// DailyEngagement is a computed model for daily metrics.
// Embedded in EventAnalytics.DailyMetrics, keyed by date string (e.g., "2024-01-15").
// Compute average time on page as TotalTimeOnPage / Views.
type DailyEngagement struct {
	Views           int32 `json:"views" bson:"views"`
	UniqueViews     int32 `json:"unique_views" bson:"unique_views"`
	RSVPs           int32 `json:"rsvps" bson:"rsvps"`
	Shares          int32 `json:"shares" bson:"shares"`
	TotalTimeOnPage int64 `json:"total_time_on_page" bson:"total_time_on_page"` // seconds
}
