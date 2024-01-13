package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
Create Post Model
*/
type PostDao struct {
	ID           primitive.ObjectID  `json:"id,omitempty" bson:"_id"`
	GroupID      primitive.ObjectID  `json:"group_id" bson:"group_id"`
	EventID      *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	Type         string              `json:"type" bson:"type"`
	Poster       string              `json:"poster,omitempty" bson:"poster,omitempty"`
	Body         string              `json:"body" bson:"body"`
	Images       *[]string           `json:"images,omitempty" bson:"images,omitempty"`
	Likes        *[]Like             `json:"likes,omitempty" bson:"likes,omitempty"`
	Comments     *[]Comment          `json:"comments,omitempty" bson:"comments,omitempty"`
	CreatedAt    *int64              `json:"created_at" bson:"created_at"`
	ExternalLink *string             `json:"external_link,omitempty" bson:"external_link,omitempty"`
}

/*
Update Post Model
*/
type PostUpdate struct {
	EventID      *primitive.ObjectID `json:"event_id,omitempty" bson:"event_id"`
	Body         *string             `json:"body" bson:"body"`
	Images       *[]string           `json:"images" bson:"images"`
	ExternalLink *string             `json:"external_link" bson:"external_link"`
}

/*
Comment
  - Comments on post
*/
type Comment struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	User      *UserSnippet       `json:"user,omitempty" bson:"user,omitempty"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

type CommentDao struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

/*
Post
  - Post objects for feed
*/
type Post struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Poster       *UserSnippet       `json:"poster,omitempty" bson:"poster,omitempty"`
	Type         string             `json:"type" bson:"type"`
	Body         string             `json:"body" bson:"body"`
	Images       *[]string          `json:"images,omitempty" bson:"images,omitempty"`
	Event        *Event             `json:"event,omitempty"`
	Likes        *[]Like            `json:"likes,omitempty" bson:"likes,omitempty"`
	Comments     *[]Comment         `json:"comments,omitempty" bson:"comments,omitempty"`
	CreatedAt    int64              `json:"created_at" bson:"created_at"`
	ExternalLink *string            `json:"external_link,omitempty" bson:"external_link,omitempty"`
}

/*
Posts Response
  - array of posts
*/
type PostsResponse struct {
	TotalPosts int    `json:"total_posts"`
	Posts      []Post `json:"posts"`
}
