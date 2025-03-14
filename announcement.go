package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AnnouncementStatus represents the possible states of an announcement
type AnnouncementStatus string

const (
	StatusPending  AnnouncementStatus = "pending"
	StatusInReview AnnouncementStatus = "in-review"
	StatusActive   AnnouncementStatus = "active"
	StatusExpired  AnnouncementStatus = "expired"
)

// AnnouncementScope defines the visibility range
type AnnouncementScope string

const (
	ScopeLocal           AnnouncementScope = "local"
	ScopeGlobal          AnnouncementScope = "global"
	ScopeApplicationWide AnnouncementScope = "application-wide"
)

// PositionConfig defines how the announcement is positioned on screen
type PositionConfig struct {
	Alignment   string `json:"alignment" bson:"alignment"`       // left, right, center
	VerticalPos string `json:"vertical_pos" bson:"vertical_pos"` // top, middle, bottom
	XOffset     int    `json:"x_offset" bson:"x_offset"`         // horizontal offset in pixels
	YOffset     int    `json:"y_offset" bson:"y_offset"`         // vertical offset in pixels
	Width       string `json:"width" bson:"width"`               // width value (%, px)
	Height      string `json:"height" bson:"height"`             // height value (%, px)
}

// TextEmphasisType defines which text element should have visual prominence
type TextEmphasisType string

const (
	EmphasisTitle    TextEmphasisType = "title"    // Title is more prominent
	EmphasisSubtitle TextEmphasisType = "subtitle" // Subtitle is more prominent
	EmphasisEqual    TextEmphasisType = "equal"    // Equal prominence
)

// TextStyleConfig defines styling options for text elements
type TextStyleConfig struct {
	FontSize      string `json:"font_size" bson:"font_size"`           // Size (px, rem, em, %)
	FontWeight    string `json:"font_weight" bson:"font_weight"`       // normal, bold, etc.
	Color         string `json:"color" bson:"color"`                   // Text color
	TextAlign     string `json:"text_align" bson:"text_align"`         // left, center, right
	LineHeight    string `json:"line_height" bson:"line_height"`       // Line height
	LetterSpacing string `json:"letter_spacing" bson:"letter_spacing"` // Letter spacing
}

// ActionButton defines properties for the interactive button
type ActionButton struct {
	Text      string `json:"text" bson:"text"`             // Button text
	URL       string `json:"url" bson:"url"`               // Target URL
	Color     string `json:"color" bson:"color"`           // Button color
	TextColor string `json:"text_color" bson:"text_color"` // Button text color
}

// Announcement represents the main announcement object
type Announcement struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Creator       UserSnippet        `json:"creator" bson:"creator"`                       // Who created the announcement
	Title         string             `json:"title" bson:"title"`                           // Main title text
	Subtitle      string             `json:"subtitle" bson:"subtitle"`                     // Secondary text
	TextEmphasis  TextEmphasisType   `json:"text_emphasis" bson:"text_emphasis"`           // Which text element has prominence
	TitleStyle    TextStyleConfig    `json:"title_style" bson:"title_style"`               // Styling for title
	SubtitleStyle TextStyleConfig    `json:"subtitle_style" bson:"subtitle_style"`         // Styling for subtitle
	MediaURL      string             `json:"media_url" bson:"media_url"`                   // Image or video URL
	MediaType     string             `json:"media_type" bson:"media_type"`                 // "image" or "video"
	ActionButton  ActionButton       `json:"action_button" bson:"action_button"`           // Call-to-action button
	Position      PositionConfig     `json:"position" bson:"position"`                     // UI positioning
	Scope         AnnouncementScope  `json:"scope" bson:"scope"`                           // Visibility scope
	Location      *GeoJSON           `json:"location,omitempty" bson:"location,omitempty"` // For local announcements
	Status        AnnouncementStatus `json:"status" bson:"status"`                         // Current state
	ActiveDate    primitive.DateTime `json:"active_date" bson:"active_date"`               // When to start showing
	ExpiryDate    primitive.DateTime `json:"expiry_date" bson:"expiry_date"`               // When to stop showing
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`                 // Creation timestamp
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"`                 // Last update timestamp
}

// AnnouncementDao is the data access object for announcements
type AnnouncementDao struct {
	Creator       *string             `json:"creator,omitempty" bson:"creator,omitempty"`
	Title         *string             `json:"title,omitempty" bson:"title,omitempty"`
	Subtitle      *string             `json:"subtitle,omitempty" bson:"subtitle,omitempty"`
	TextEmphasis  *TextEmphasisType   `json:"text_emphasis,omitempty" bson:"text_emphasis,omitempty"`
	TitleStyle    *TextStyleConfig    `json:"title_style,omitempty" bson:"title_style,omitempty"`
	SubtitleStyle *TextStyleConfig    `json:"subtitle_style,omitempty" bson:"subtitle_style,omitempty"`
	MediaURL      *string             `json:"media_url,omitempty" bson:"media_url,omitempty"`
	MediaType     *string             `json:"media_type,omitempty" bson:"media_type,omitempty"`
	ActionButton  *ActionButton       `json:"action_button,omitempty" bson:"action_button,omitempty"`
	Position      *PositionConfig     `json:"position,omitempty" bson:"position,omitempty"`
	Scope         *AnnouncementScope  `json:"scope,omitempty" bson:"scope,omitempty"`
	Location      *GeoJSON            `json:"location,omitempty" bson:"location,omitempty"`
	Status        *AnnouncementStatus `json:"status,omitempty" bson:"status,omitempty"`
	ActiveDate    *primitive.DateTime `json:"active_date,omitempty" bson:"active_date,omitempty"`
	ExpiryDate    *primitive.DateTime `json:"expiry_date,omitempty" bson:"expiry_date,omitempty"`
	UpdatedAt     *primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt     *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// AnnouncementsResponse wraps a list of announcements for API responses
type AnnouncementsResponse struct {
	TotalAnnouncements int            `json:"total_announcements"`
	Announcements      []Announcement `json:"announcements"`
}
