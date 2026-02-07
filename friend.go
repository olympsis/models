package models

import "go.mongodb.org/mongo-driver/v2/bson"

/*
Friend
  - Friend object
*/
type Friend struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	CreatedAt bson.DateTime `json:"created_at" bson:"created_at"`
}

/*
Friend Request

  - friend request object
*/
type FriendRequest struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Requestor string             `json:"requestor" bson:"requestor"`
	Requestee string             `json:"requestee" bson:"requestee"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt bson.DateTime `json:"created_at" bson:"created_at"`
}

/*
Friend Requests

  - total number of friend request

  - friend requests
*/
type FriendRequests struct {
	TotalRequests int             `json:"total_requests"`
	Requests      []FriendRequest `json:"requests"`
}
