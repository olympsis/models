package models

import (
	"errors"

	webpush "github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/v2/bson"
)

/*
Notification represents a basic notification message
*/
type Notification struct {
	Title    string                 `json:"title"`
	Body     string                 `json:"body"`
	Topic    string                 `json:"topic"`
	Priority int                    `json:"priority"`
	Data     map[string]interface{} `json:"data"`
}

/*
NotificationPreference stores user notification settings
*/
type NotificationPreference struct {
	Types      map[string]bool `json:"types" bson:"types"`           // push, email, phone
	Categories map[string]bool `json:"categories" bson:"categories"` // groups, events
	UpdatedAt  bson.DateTime   `json:"updated_at" bson:"updated_at"`
}

/*
NotificationDevice represents a user's device for push notifications
*/
type NotificationDevice struct {
	DeviceID        string                `json:"device_id,omitempty" bson:"device_id,omitempty"`
	Token           string                `json:"token,omitempty" bson:"token,omitempty"`
	WebSubscription *webpush.Subscription `json:"web_subscription,omitempty" bson:"web_subscription,omitempty"`
	DeviceInfo      DeviceInfo            `json:"device_info" bson:"device_info"`
	Active          *bool                 `json:"active,omitempty" bson:"active,omitempty"`
	CreatedAt       *bson.DateTime        `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       *bson.DateTime        `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

/*
NotificationLog tracks notification delivery status
*/
type NotificationLog struct {
	ID             bson.ObjectID `json:"id,omitempty" bson:"_id"`
	NotificationID bson.ObjectID `json:"notification_id" bson:"notification_id"`
	Platform       string        `json:"platform" bson:"platform"`
	Status         string        `json:"status" bson:"status"`
	Error          *string       `json:"error,omitempty" bson:"error,omitempty"`
	CreatedAt      bson.DateTime `json:"created_at" bson:"created_at"`
}

/*
PushNotification represents the core notification content
*/
type PushNotification struct {
	ID        bson.ObjectID          `json:"id,omitempty" bson:"_id"`
	Title     string                 `json:"title" bson:"title"`
	Body      string                 `json:"body" bson:"body"`
	Type      string                 `json:"type" bson:"type"`         // push, email, sms
	Category  string                 `json:"category" bson:"category"` // groups, events, announcements
	Data      map[string]interface{} `json:"data" bson:"data"`
	ExpiresAt *bson.DateTime         `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	CreatedAt bson.DateTime          `json:"created_at" bson:"created_at"`
}

/*
UserNotification represents a notification instance for a specific user
*/
type UserNotification struct {
	ID             bson.ObjectID  `json:"id,omitempty" bson:"_id"`
	UserID         string         `json:"user_id" bson:"user_id"`
	NotificationID bson.ObjectID  `json:"notification_id" bson:"notification_id"`
	IsRead         bool           `json:"is_read" bson:"is_read"`
	IsArchived     bool           `json:"is_archived" bson:"is_archived"`
	ReadAt         *bson.DateTime `json:"read_at,omitempty" bson:"read_at,omitempty"`
	CreatedAt      bson.DateTime  `json:"created_at" bson:"created_at"`
}

/*
NotificationTopic represents a topic that users can subscribe to
*/
type NotificationTopic struct {
	ID        bson.ObjectID  `json:"id,omitempty" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	Type      string         `json:"type" bson:"type"` //groups, events
	Users     []string       `json:"users" bson:"users"`
	IsActive  bool           `json:"is_active" bson:"is_active"`
	CreatedAt bson.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt *bson.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

/*
NotificationTopicDao for database operations
*/
type NotificationTopicDao struct {
	ID        *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      *string        `json:"name,omitempty" bson:"name,omitempty"`
	Type      *string        `json:"type,omitempty" bson:"type,omitempty"`
	Users     *[]string      `json:"users,omitempty" bson:"users,omitempty"`
	IsActive  *bool          `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *bson.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

/*
NotificationListResponse wraps the response for notification list queries
*/
type NotificationListResponse struct {
	TotalNotifications int                `json:"total_notifications"`
	UnreadCount        int                `json:"unread_count"`
	Notifications      []NotificationItem `json:"notifications"`
}

/*
NotificationItem combines PushNotification and user-specific state
*/
type NotificationItem struct {
	ID        bson.ObjectID          `json:"id" bson:"_id"`
	Title     string                 `json:"title" bson:"title"`
	Body      string                 `json:"body" bson:"body"`
	Type      string                 `json:"type" bson:"type"`
	Category  string                 `json:"category" bson:"category"`
	Data      map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"`
	IsRead    bool                   `json:"is_read" bson:"is_read"`
	CreatedAt bson.DateTime          `json:"created_at" bson:"created_at"`
}

/*
NotificationUpdateRequest for marking notifications as read/archived
*/
type NotificationUpdateRequest struct {
	NotificationIDs []string `json:"notification_ids"`
	Action          string   `json:"action"` // read, unread, archive
}

/*
NotificationTopicUpdateRequest for topic modifications
*/
type NotificationTopicUpdateRequest struct {
	Action string   `json:"action" bson:"action"` // subscribe, unsubscribe, update
	Users  []string `json:"users" bson:"users"`
}

/*
Notification Push request. This helps us send pushes to topics or a list of users
*/
type NotificationPushRequest struct {
	Topic        *string          `json:"topic,omitempty" bson:"topic,omitempty"`
	Users        *[]string        `json:"users,omitempty" bson:"users,omitempty"`
	Notification PushNotification `json:"notification" bson:"notification"`
}

func (n *NotificationPushRequest) Validate() error {
	if n.Topic == nil && n.Users == nil {
		return errors.New("a push request requires either a topic or a list of users")
	}
	return nil
}

type NotificationPayload struct {
	Type    NotificationType `json:"type"`
	ClubID  bson.ObjectID    `json:"club_id,omitempty"`
	EventID bson.ObjectID    `json:"event_id,omitempty"`
	PostID  bson.ObjectID    `json:"post_id,omitempty"`
}
