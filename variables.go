package models

// User genders
type Gender string

const (
	GenderMale           Gender = "male"
	GenderFemale         Gender = "female"
	GenderNonBinary      Gender = "non_binary"
	GenderPreferNotToSay Gender = "prefer_not_to_say"
)

// Club member roles
type MemberRole string

const (
	OwnerMember  MemberRole = "OWNER"
	AdminMember  MemberRole = "ADMIN"
	MemberMember MemberRole = "MEMBER"
)

// Media types
type MediaType string

const (
	ImageMediaType MediaType = "IMAGE"
	VideoMediaType MediaType = "VIDEO"
)

// Visibility Scope represents the possible visibility states of many models
type VisibilityScope string

const (
	PublicVisibilityScope  VisibilityScope = "PUBLIC"
	PrivateVisibilityScope VisibilityScope = "PRIVATE"
	GroupVisibilityScope   VisibilityScope = "GROUP"
)

// Application status:  pending | accepted | denied
type ApplicationStatus string

const (
	PendingApplicationStatus  ApplicationStatus = "PENDING"
	AcceptedApplicationStatus ApplicationStatus = "ACCEPTED"
	DeniedApplicationStatus   ApplicationStatus = "DENIED"
)

// Competition state: pending | live | completed
type CompetitionState string

const (
	PendingCompetitionState   CompetitionState = "PENDING"
	LiveCompetitionState      CompetitionState = "LIVE"
	CompletedCompetitionState CompetitionState = "COMPLETED"
)

// RSVP status represents the possible status of an rsvp
type RSVPStatus int

const (
	RSVPMaybe    RSVPStatus = 0
	RSVPYes      RSVPStatus = 1
	RSVPWaitlist RSVPStatus = 2
)

type ClubFinancialAccountStatus string

const (
	ClubFinancialAccountStatusPending    = "PENDING"
	ClubFinancialAccountStatusActive     = "ACTIVE"
	ClubFinancialAccountStatusRestricted = "RESTRICTED"
)

// Event Competition formats
type CompetitionFormats string

const (
	// General Team Sports (Soccer, Basketball, Volleyball, Football, Flag Football, Padel, Pickleball, Badminton, Ping-Pong, Racketball)
	Bracket           CompetitionFormats = "bracket"            // Knockout-style tournament
	League            CompetitionFormats = "league"             // Regular season format
	RoundRobin        CompetitionFormats = "round_robin"        // Each team plays all other teams
	SingleElimination CompetitionFormats = "single_elimination" // One loss and you're out
	DoubleElimination CompetitionFormats = "double_elimination" // Two losses before elimination
	BestOf3           CompetitionFormats = "best_of_3"          // First to win 2 games
	BestOf5           CompetitionFormats = "best_of_5"          // First to win 3 games
	WinnerStaysOn     CompetitionFormats = "winner_stays_on"    // Winners keep playing, losers rotate out

	// Team Formats (Soccer, Basketball, Volleyball, Football, Flag Football, Padel, Pickleball, Badminton, Ping-Pong, Racketball)
	Versus2  CompetitionFormats = "2v2"
	Versus3  CompetitionFormats = "3v3"
	Versus4  CompetitionFormats = "4v4"
	Versus5  CompetitionFormats = "5v5"
	Versus6  CompetitionFormats = "6v6"
	Versus7  CompetitionFormats = "7v7"
	Versus8  CompetitionFormats = "8v8"
	Versus9  CompetitionFormats = "9v9"
	Versus10 CompetitionFormats = "10v10"
	Versus11 CompetitionFormats = "11v11"

	// Individual Sports (Running, Cycling)
	TimeTrial CompetitionFormats = "time_trial" // Athletes race against the clock

	// Running
	Sprint       CompetitionFormats = "sprint"        // Short-distance race (e.g., 100m, 200m)
	LongDistance CompetitionFormats = "long_distance" // Longer races (e.g., 5K, 10K, marathon)
	Relay        CompetitionFormats = "relay"         // Team race with baton passing

	// Cycling
	RoadRace  CompetitionFormats = "road_race"  // Mass-start long-distance race
	Criterium CompetitionFormats = "criterium"  // Short circuit, multiple laps
	StageRace CompetitionFormats = "stage_race" // Multi-day competition (e.g., Tour de France)

	// Golf
	StrokePlay         CompetitionFormats = "stroke_play"         // Total strokes over the round(s) determine the winner
	MatchPlay          CompetitionFormats = "match_play"          // Head-to-head format, winning holes instead of strokes
	Scramble           CompetitionFormats = "scramble"            // Teams play the best shot among their members
	BestBall           CompetitionFormats = "best_ball"           // Each player plays their ball, best score counts for the team
	Stableford         CompetitionFormats = "stableford"          // Points awarded based on score per hole
	SkinsGame          CompetitionFormats = "skins_game"          // Each hole has a prize (skin), won outright by lowest score
	AlternateShot      CompetitionFormats = "alternate_shot"      // Two-player teams alternate shots on the same ball
	Shamble            CompetitionFormats = "shamble"             // Similar to scramble but players play from the best tee shot
	ModifiedStableford CompetitionFormats = "modified_stableford" // Variation of Stableford with adjusted point values
	Scratch            CompetitionFormats = "scratch"             // No handicaps, raw stroke count matters

	// Climbing
	Bouldering    CompetitionFormats = "bouldering"     // Short, difficult climbing routes, no ropes
	LeadClimbing  CompetitionFormats = "lead_climbing"  // Climbing as high as possible on a tall wall
	SpeedClimbing CompetitionFormats = "speed_climbing" // Race to the top
)

type NotificationType string

const (
	NewClubApplicationType           NotificationType = "new_club_application"
	ClubApplicationUpdateType        NotificationType = "club_application_update"
	RankingChangeType                NotificationType = "ranking_change"
	SuspensionType                   NotificationType = "suspension"
	ExpulsionType                    NotificationType = "expulsion"
	PostingApprovalRequestType       NotificationType = "posting_approval_request"
	PostingApprovalRequestUpdateType NotificationType = "posting_approval_request_update"

	MemberReportType NotificationType = "member_report"

	NewPostType       NotificationType = "new_post"
	PostReportType    NotificationType = "post_report"
	CommentReportType NotificationType = "comment_report"

	EventReminderType      NotificationType = "event_reminder"
	DailyEventSummaryType  NotificationType = "daily_event_summary"
	WeeklyEventSummaryType NotificationType = "weekly_event_summary"

	NewEventType                        NotificationType = "new_event"
	NewEventCommentType                 NotificationType = "new_event_comment"
	EventParticipantUpdateType          NotificationType = "event_participant_update"
	EventParticipantKickType            NotificationType = "event_participant_kick"
	EventParticipantWaitlistUpgradeType NotificationType = "event_participant_waitlist_upgrade"
	EventCancellation                   NotificationType = "event_cancellation"

	NewAnnouncementType NotificationType = "new_announcement"
)

type NotificationTopicType string

const (
	EventNotificationTopic NotificationTopicType = "event"

	ClubNotificationTopic         NotificationTopicType = "club"
	OrganizationNotificationTopic NotificationTopicType = "organization"

	PostNotificationTopic         NotificationTopicType = "post"
	AnnouncementNotificationTopic NotificationTopicType = "announcement"

	MessagesNotificationTopic NotificationTopicType = "messages"
)

type RecurrencePattern string

const (
	RecurrencePatternDaily   RecurrencePattern = "DAILY"
	RecurrencePatternWeekly  RecurrencePattern = "WEEKLY"
	RecurrencePatternMonthly RecurrencePattern = "MONTHLY"
)

type UserType string

const (
	MemberUserType     UserType = "MEMBER"
	NonMemberUserType  UserType = "NON_MEMBER"
	AnonMemberUserType UserType = "ANON_MEMBER"
)

// Event types
type EventType string

const (
	RegularEventType    EventType = "REGULAR"
	LeagueEventType     EventType = "LEAGUE"
	TournamentEventType EventType = "TOURNAMENT"
	ClassEventType      EventType = "CLASS"
)

type OrganizerType string

const (
	OrganizerTypeGroup        OrganizerType = "GROUP"
	OrganizerTypeOrganization OrganizerType = "ORGANIZATION"
)
