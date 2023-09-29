package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Avatar      string             `json:"avatar" bson:"avatar"`
	Videos      []*Video           `json:"videos" bson:"videos"`
}

type Video struct {
	URI   string `json:"uri"  bson:"uri"`
	Title string `json:"title"  bson:"title"`
}
