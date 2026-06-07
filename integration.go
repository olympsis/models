package models

import "go.mongodb.org/mongo-driver/v2/bson"

// ChatPlatform identifies a third-party chat platform a club can integrate with.
//
// WhatsApp was intentionally excluded: its Business Cloud API does not support group
// chat and group automation requires ToS-violating unofficial clients. Telegram and
// Discord both expose official, ToS-compliant bot APIs that support group participation,
// message reading, polls, and buttons.
type ChatPlatform string

const (
	PlatformTelegram ChatPlatform = "telegram"
	PlatformDiscord  ChatPlatform = "discord"
)

// ChatLinkStatus is the lifecycle state of a ClubChatLink.
type ChatLinkStatus string

const (
	ChatLinkPending ChatLinkStatus = "pending" // code issued, awaiting /link confirmation in the chat
	ChatLinkActive  ChatLinkStatus = "active"  // chat confirmed and bound to the club
)

// ClubChatLink binds an Olympsis club to a Telegram group or Discord channel that the
// "bot father" posts event reminders to and reads RSVPs from. It follows the
// separate-collection-per-club pattern established by ClubFinancialAccount.
type ClubChatLink struct {
	ID       bson.ObjectID  `json:"id" bson:"_id"`
	ClubID   bson.ObjectID  `json:"club_id" bson:"club_id"`
	Platform ChatPlatform   `json:"platform" bson:"platform"`
	Status   ChatLinkStatus `json:"status" bson:"status"`
	LinkCode string         `json:"link_code,omitempty" bson:"link_code,omitempty"` // short-lived code an admin runs as /link <code>

	// Telegram uses ChatID (group/supergroup id). Discord uses GuildID + ChannelID.
	ChatID    string `json:"chat_id,omitempty" bson:"chat_id,omitempty"`
	GuildID   string `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	ChannelID string `json:"channel_id,omitempty" bson:"channel_id,omitempty"`

	Title    string `json:"title,omitempty" bson:"title,omitempty"`         // chat/channel name captured at link time
	LinkedBy string `json:"linked_by,omitempty" bson:"linked_by,omitempty"` // user_id of the admin who initiated

	CreatedAt bson.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt *bson.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// ClubChatLinkDao is the partial-update / insert variant of ClubChatLink.
type ClubChatLinkDao struct {
	ID        *bson.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	ClubID    *bson.ObjectID  `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Platform  *ChatPlatform   `json:"platform,omitempty" bson:"platform,omitempty"`
	Status    *ChatLinkStatus `json:"status,omitempty" bson:"status,omitempty"`
	LinkCode  *string         `json:"link_code,omitempty" bson:"link_code,omitempty"`
	ChatID    *string         `json:"chat_id,omitempty" bson:"chat_id,omitempty"`
	GuildID   *string         `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	ChannelID *string         `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	Title     *string         `json:"title,omitempty" bson:"title,omitempty"`
	LinkedBy  *string         `json:"linked_by,omitempty" bson:"linked_by,omitempty"`
	CreatedAt *bson.DateTime  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *bson.DateTime  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// ChatIdentity maps a platform-native user (Telegram/Discord) to an Olympsis user so
// bot-captured RSVPs can be attributed to the right member. UserID is empty until the
// sender is resolved (e.g. via roster name matching or an explicit linking command).
type ChatIdentity struct {
	ID             bson.ObjectID `json:"id" bson:"_id"`
	Platform       ChatPlatform  `json:"platform" bson:"platform"`
	PlatformUserID string        `json:"platform_user_id" bson:"platform_user_id"`
	UserID         string        `json:"user_id,omitempty" bson:"user_id,omitempty"`
	DisplayName    string        `json:"display_name,omitempty" bson:"display_name,omitempty"`
	Handle         string        `json:"handle,omitempty" bson:"handle,omitempty"`
	CreatedAt      bson.DateTime `json:"created_at" bson:"created_at"`
}

// IntegrationLinkResponse is returned to a club admin when they start linking a chat.
// It carries the code the admin must run in-chat plus platform setup instructions.
type IntegrationLinkResponse struct {
	Platform     ChatPlatform  `json:"platform"`
	LinkCode     string        `json:"link_code"`
	Instructions string        `json:"instructions"`
	InviteURL    string        `json:"invite_url,omitempty"` // Discord bot OAuth2 invite URL
	ExpiresAt    bson.DateTime `json:"expires_at"`
}

// IntegrationConfirmRequest is sent by the bot service to the main server when an admin
// runs /link <code> inside the target chat, completing the binding.
type IntegrationConfirmRequest struct {
	Platform  ChatPlatform `json:"platform"`
	LinkCode  string       `json:"link_code"`
	ChatID    string       `json:"chat_id,omitempty"`
	GuildID   string       `json:"guild_id,omitempty"`
	ChannelID string       `json:"channel_id,omitempty"`
	Title     string       `json:"title,omitempty"`
}

// ChatRosterMember is a lightweight club-member descriptor used for RSVP attribution.
type ChatRosterMember struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name,omitempty"`
}

// BotReminderRequest is sent by the main server to the bot service to post an event
// reminder into a linked chat. The roster lets the LLM map free-text replies back to
// Olympsis members.
type BotReminderRequest struct {
	EventID   string             `json:"event_id"`
	ClubID    string             `json:"club_id"`
	Platform  ChatPlatform       `json:"platform"`
	ChatID    string             `json:"chat_id,omitempty"`
	GuildID   string             `json:"guild_id,omitempty"`
	ChannelID string             `json:"channel_id,omitempty"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	StartsAt  bson.DateTime      `json:"starts_at"`
	Roster    []ChatRosterMember `json:"roster,omitempty"`
}

// BotRSVPReport is sent by the bot service back to the main server after the LLM parses
// replies in a linked chat into structured RSVP intent.
type BotRSVPReport struct {
	EventID  string          `json:"event_id"`
	Platform ChatPlatform    `json:"platform"`
	Results  []BotRSVPResult `json:"results"`
}

// BotRSVPResult is a single parsed RSVP outcome. Olympsis has no explicit "not coming"
// RSVP status (see RSVPStatus), so declines are represented with the Decline flag and
// handled by removing the participant.
type BotRSVPResult struct {
	UserID         string     `json:"user_id,omitempty"`          // resolved Olympsis user_id (empty if unmatched)
	PlatformUserID string     `json:"platform_user_id,omitempty"` // raw sender id for later mapping
	DisplayName    string     `json:"display_name,omitempty"`
	Status         RSVPStatus `json:"status"`     // RSVPYes / RSVPMaybe
	Decline        bool       `json:"decline"`    // true => "not coming" (remove participant)
	Confidence     float64    `json:"confidence"` // 0..1 from the LLM
}

// BotParticipantRequest is the body for the internal bot-driven RSVP endpoint on the
// main server (POST /v1/events/{id}/participants/bot). The server records the RSVP on
// behalf of a resolved user.
type BotParticipantRequest struct {
	UserID  string       `json:"user_id"`
	Status  RSVPStatus   `json:"status"`
	Decline bool         `json:"decline"`
	Source  ChatPlatform `json:"source"`
}
