package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Post struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	Poster       *UserSnippet       `json:"poster,omitempty" bson:"poster,omitempty"`
	Type         string             `json:"type" bson:"type"`
	Body         string             `json:"body" bson:"body"`
	Images       []string           `json:"images" bson:"images"`
	Event        *Event             `json:"event,omitempty"`
	Reactions    []Reaction         `json:"reactions" bson:"reactions"`
	Comments     []PostComment      `json:"comments" bson:"comments"`
	ExternalLink *string            `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  bool               `json:"is_sensitive" bson:"is_sensitive"`
	CreatedAt    bson.DateTime `json:"created_at" bson:"created_at"`
}

type PostDao struct {
	ID           *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GroupID      *bson.ObjectID `json:"group_id" bson:"group_id,omitempty"`
	EventID      *bson.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	PostID       *bson.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Type         *string             `json:"type,omitempty" bson:"type,omitempty"`
	Poster       *string             `json:"poster,omitempty" bson:"poster,omitempty"`
	Body         *string             `json:"body,omitempty" bson:"body,omitempty"`
	Images       *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	ExternalLink *string             `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  *bool               `json:"is_sensitive,omitempty" bson:"is_sensitive,omitempty"`
	CreatedAt    *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type PostComment struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt bson.DateTime `json:"created_at" bson:"created_at"`
}

type PostCommentDao struct {
	ID        *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    *string             `json:"user_id,omitempty" bson:"user_id.omitempty"`
	PostID    *bson.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Text      *string             `json:"text,omitempty" bson:"text,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at.omitempty"`
}

type PostsResponse struct {
	TotalPosts int    `json:"total_posts"`
	Posts      []Post `json:"posts"`
}
