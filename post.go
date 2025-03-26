package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Poster       *UserSnippet       `json:"poster,omitempty" bson:"poster,omitempty"`
	Type         string             `json:"type" bson:"type"`
	Body         string             `json:"body" bson:"body"`
	Images       []string           `json:"images" bson:"images"`
	Event        *Event             `json:"event,omitempty"`
	Likes        []Like             `json:"likes" bson:"likes"`
	Comments     []Comment          `json:"comments" bson:"comments"`
	ExternalLink *string            `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  bool               `json:"is_sensitive" bson:"is_sensitive"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
}

type PostDao struct {
	ID           *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GroupID      *primitive.ObjectID `json:"group_id" bson:"group_id,omitempty"`
	EventID      *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	PostID       *primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Type         *string             `json:"type,omitempty" bson:"type,omitempty"`
	Poster       *string             `json:"poster,omitempty" bson:"poster,omitempty"`
	Body         *string             `json:"body,omitempty" bson:"body,omitempty"`
	Images       *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	ExternalLink *string             `json:"external_link,omitempty" bson:"external_link,omitempty"`
	IsSensitive  *bool               `json:"is_sensitive,omitempty" bson:"is_sensitive,omitempty"`
	CreatedAt    *primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

type CommentDao struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    *string             `json:"user_id,omitempty" bson:"user_id.omitempty"`
	PostID    *primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Text      *string             `json:"text,omitempty" bson:"text,omitempty"`
	CreatedAt *primitive.DateTime `json:"created_at,omitempty" bson:"created_at.omitempty"`
}

type PostsResponse struct {
	TotalPosts int    `json:"total_posts"`
	Posts      []Post `json:"posts"`
}
