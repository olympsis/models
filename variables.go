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
	OwnerMember  MemberRole = "owner"
	AdminMember  MemberRole = "admin"
	MemberMember MemberRole = "member"
)

// Media types
type MediaType string

const (
	ImageMediaType MediaType = "image"
	VideoMediaType MediaType = "video"
)

// Visibility Scope represents the possible visibility states of many models
type VisibilityScope int

const (
	PublicVisibilityScope  VisibilityScope = 0
	PrivateVisibilityScope VisibilityScope = 1
	GroupVisibilityScope   VisibilityScope = 2
)

// Application status:  pending | accepted | denied
type ApplicationStatus string

const (
	PendingApplicationStatus  = "pending"
	AcceptedApplicationStatus = "accepted"
	DeniedApplicationStatus   = "denied"
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
	ClubFinancialAccountStatusPending    = "pending"
	ClubFinancialAccountStatusActive     = "active"
	ClubFinancialAccountStatusRestricted = "restricted"
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
	NewClubApplication               NotificationType = "new_club_application"
	ClubApplicationUpdate            NotificationType = "club_application_update"
	ClubRankingChange                NotificationType = "club_ranking_change"
	ClubSuspension                   NotificationType = "club_suspension"
	ClubExpulsion                    NotificationType = "club_expulsion"
	ClubPostingApprovalRequest       NotificationType = "club_posting_approval_request"
	ClubPostingApprovalRequestUpdate NotificationType = "club_posting_approval_request_update"
	ClubNewPost                      NotificationType = "club_new_post"
	ClubPostReport                   NotificationType = "club_post_report"
	ClubCommentReport                NotificationType = "club_comment_report"
	ClubMemberReport                 NotificationType = "club_member_report"
	ClubNewEvent                     NotificationType = "club_new_event"
	NewEventComment                  NotificationType = "new_event_comment"
	EventParticipantUpdate           NotificationType = "event_participant_update"
	EventReminder                    NotificationType = "event_reminder"
	DailyEventSummary                NotificationType = "daily_event_summary"
	WeeklyEventSummary               NotificationType = "weekly_event_summary"
)
