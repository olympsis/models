package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Poster *UserSnippet       `json:"poster,omitempty" bson:"poster,omitempty"`

	Organizers []Organizer       `json:"organizers" bson:"organizers"`
	Sponsors   []Organization    `json:"sponsors" bson:"sponsors"`
	Venues     []VenueDescriptor `json:"venues" bson:"venues"`

	MediaURL  string    `json:"media_url" bson:"media_url"`
	MediaType MediaType `json:"media_type" bson:"media_type"`

	Title  string   `json:"title" bson:"title"`
	Body   string   `json:"body" bson:"body"`
	Tags   []string `json:"tags" bson:"tags"`
	Sports []string `json:"sports" bson:"sports"`

	Config       *EventConfig       `json:"config,omitempty" bson:"config,omitempty"`
	FormatConfig *EventFormatConfig `json:"format_config,omitempty" bson:"format_config,omitempty"`

	StartTime primitive.DateTime `json:"start_time" bson:"start_time"`
	StopTime  primitive.DateTime `json:"stop_time" bson:"stop_time"`

	Participants         *[]Participant      `json:"participants,omitempty" bson:"participants,omitempty"`
	ParticipantsWaitlist *[]Participant      `json:"participants_waitlist,omitempty" bson:"participants_waitlist,omitempty"`
	ParticipantsConfig   *ParticipantsConfig `json:"participants_config,omitempty" bson:"participants_config,omitempty"`

	Teams         *[]Team      `json:"teams,omitempty" bson:"teams,omitempty"`
	TeamsWaitlist *[]Team      `json:"teams_waitlist,omitempty" bson:"teams_waitlist,omitempty"`
	TeamsConfig   *TeamsConfig `json:"teams_config,omitempty" bson:"teams_config,omitempty"`

	Comments []EventComment `json:"comments" bson:"comments"`

	Visibility   VisibilityScope `json:"visibility" bson:"visibility"`
	ExternalLink *EventLink      `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  bool            `json:"is_sensitive" bson:"is_sensitive"`

	CreatedAt   primitive.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt   *primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CancelledAt *primitive.DateTime `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`
	ArchivedAt  *primitive.DateTime `json:"archived_at,omitempty" bson:"archived_at,omitempty"`

	RecurrenceConfig *EventRecurrenceConfig `json:"recurrence_config,omitempty" bson:"recurrence_config,omitempty"`
}

type EventDao struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PosterID *string             `json:"poster_id,omitempty" bson:"poster_id,omitempty"`

	Organizers *[]Organizer          `json:"organizers,omitempty" bson:"organizers,omitempty"`
	Sponsors   *[]primitive.ObjectID `json:"sponsors,omitempty" bson:"sponsors,omitempty"`
	Venues     *[]VenueDescriptor    `json:"venues,omitempty" bson:"venues,omitempty"`

	// Event media
	MediaURL  *string    `json:"media_url,omitempty" bson:"media_url,omitempty"`
	MediaType *MediaType `json:"media_type,omitempty" bson:"media_type,omitempty"`

	// Event information
	Title  *string   `json:"title,omitempty" bson:"title,omitempty"`
	Body   *string   `json:"body,omitempty" bson:"body,omitempty"`
	Sports *[]string `json:"sports,omitempty" bson:"sports,omitempty"`
	Tags   *[]string `json:"tags,omitempty" bson:"tags,omitempty"`

	// Event configuration
	Config       *EventConfig       `json:"config,omitempty" bson:"config,omitempty"`
	FormatConfig *EventFormatConfig `json:"format_config,omitempty" bson:"format_config,omitempty"`

	StartTime *primitive.DateTime `json:"start_time,omitempty" bson:"start_time,omitempty"`
	StopTime  *primitive.DateTime `json:"stop_time,omitempty" bson:"stop_time,omitempty"`

	ParticipantsConfig *ParticipantsConfig `json:"participants_config,omitempty" bson:"participants_config,omitempty"`
	TeamsConfig        *TeamsConfig        `json:"teams_config,omitempty" bson:"teams_config,omitempty"`

	Visibility   *VisibilityScope `json:"visibility,omitempty" bson:"visibility,omitempty"`
	ExternalLink *EventLink       `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  *bool            `json:"is_sensitive,omitempty" bson:"is_sensitive,omitempty"`

	CreatedAt   *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CancelledAt *primitive.DateTime `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`
	ArchivedAt  *primitive.DateTime `json:"archived_at,omitempty" bson:"archived_at,omitempty"`

	RecurrenceConfig *EventRecurrenceConfig `json:"recurrence_config,omitempty" bson:"recurrence_config,omitempty"`
}

type EventConfig struct {
	HidePoster   *bool `json:"hide_poster,omitempty" bson:"hide_poster,omitempty"`
	HideLocation *bool `json:"hide_location,omitempty" bson:"hide_location,omitempty"`
}

type EventFormatConfig struct {
	// Type identifiers
	IsCompetition       *bool               `json:"is_competition,omitempty" bson:"is_competition,omitempty"`
	IsCompetitionGame   *bool               `json:"is_competition_game,omitempty" bson:"is_competition_game,omitempty"`
	ParentCompetitionID *primitive.ObjectID `json:"parent_competition_id,omitempty" bson:"parent_competition_id,omitempty"`
	CompetitionState    *CompetitionState   `json:"competition_state,omitempty" bson:"competition_state,omitempty"`

	// Format details
	Formats      *[]string `json:"formats,omitempty" bson:"formats,omitempty"`
	Rounds       *int32    `json:"rounds,omitempty" bson:"rounds,omitempty"`
	CurrentRound *int32    `json:"current_round,omitempty" bson:"current_round,omitempty"`
	BracketData  *any      `json:"bracket_data,omitempty" bson:"bracket_data,omitempty"`

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
	TotalEvents int32   `json:"total_events"`
	Events      []Event `json:"events"`
}

/*********************
* PARTICIPANT MODELS *
*********************/
type Participant struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	User        *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Status      RSVPStatus         `json:"status" bson:"status"`
	EventID     primitive.ObjectID `json:"event_id" bson:"event_id"`
	IsAnonymous bool               `json:"is_anonymous" bson:"is_anonymous"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
}

type ParticipantDao struct {
	ID          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      *string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Status      *RSVPStatus         `json:"status,omitempty" bson:"status,omitempty"`
	EventID     *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	IsAnonymous *bool               `json:"is_anonymous,omitempty" bson:"is_anonymous,omitempty"`
	CreatedAt   *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type ParticipantsConfig struct {
	HasWaitlist      *bool  `json:"has_waitlist,omitempty" bson:"has_waitlist,omitempty"`
	MinParticipants  *int32 `json:"min_participants,omitempty" bson:"min_participants,omitempty"`
	MaxParticipants  *int32 `json:"max_participants,omitempty" bson:"max_participants,omitempty"`
	HideParticipants *bool  `json:"hide_participants,omitempty" bson:"hide_participants,omitempty"`
}

/**************
* TEAM MODELS *
**************/
type Team struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Members   []TeamMember       `json:"members" bson:"members"`
	EventID   primitive.ObjectID `json:"event_id" bson:"event_id"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

type TeamDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      *string             `json:"name,omitempty" bson:"name,omitempty"`
	Members   *[]TeamMemberDao    `json:"members,omitempty" bson:"members,omitempty"`
	EventID   *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type TeamMember struct {
	User        *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	IsAnonymous bool               `json:"is_anonymous" bson:"is_anonymous"`
	JoinedAt    primitive.DateTime `json:"joined_at" bson:"joined_at"`
}

type TeamMemberDao struct {
	UserID      *string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	IsAnonymous *bool               `json:"is_anonymous,omitempty" bson:"is_anonymous,omitempty"`
	JoinedAt    *primitive.DateTime `json:"joined_at,omitempty" bson:"joined_at,omitempty"`
}

type TeamsConfig struct {
	HasWaitlist *bool  `json:"has_waitlist,omitempty" bson:"has_waitlist,omitempty"`
	MinTeams    *int32 `json:"min_teams,omitempty" bson:"min_teams,omitempty"`
	MaxTeams    *int32 `json:"max_teams,omitempty" bson:"max_teams,omitempty"`
	MaxTeamSize *int32 `json:"max_team_size,omitempty" bson:"max_team_size,omitempty"`
	HideTeams   *bool  `json:"hide_teams,omitempty" bson:"hide_teams,omitempty"`
}

/***************
* OTHER MODELS *
***************/

type FullEventError struct {
	MSG   string `json:"msg"`
	Event Event  `json:"event"`
}

type Organizer struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type OrganizerType      `json:"type" bson:"type"`
}

type NewEventDao struct {
	Event       EventDao           `json:"event"`
	IncludeHost *bool              `json:"include_host,omitempty"`
	Recurrence  *RecurrenceOptions `json:"recurrence,omitempty"`
}

type RecurrenceOptions struct {
	Pattern  RecurrencePattern  `json:"pattern"`  // "DAILY", "WEEKLY", "MONTHLY"
	EndTime  primitive.DateTime `json:"end_time"` // Must be after the event's start_time
	Interval int                `json:"interval"` // e.g., every 2 weeks
}

type EventComment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Text      string             `json:"text" bson:"text"`
	EventID   primitive.ObjectID `json:"event_id" bson:"event_id"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

type EventCommentDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    *string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Text      *string             `json:"text,omitempty" bson:"text,omitempty"`
	EventID   *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type EventLink struct {
	Title string `json:"title" bson:"title"`
	URL   string `json:"url" bson:"url"`
}
