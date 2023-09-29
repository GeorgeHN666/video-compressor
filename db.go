package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	URI      string
	DATABASE string
)

type DB struct {
	C *mongo.Client
}

func StartDB() *DB {

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return nil
	}

	err = c.Ping(context.TODO(), nil)
	if err != nil {
		return nil
	}

	return &DB{
		C: c,
	}
}

func (s *DB) InsertPost(p *Post) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.C.Database(DATABASE).Collection("post")

	p.ID = primitive.NewObjectID()

	_, err := db.InsertOne(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) GetPost(id string) (*Post, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.C.Database(DATABASE).Collection("post")

	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": i,
	}

	var res Post
	err = db.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *DB) GetPosts() ([]*Post, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.C.Database(DATABASE).Collection("post")

	var res []*Post

	cursor, err := db.Find(ctx, bson.M{})
	if err != nil {
		cursor.Close(ctx)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var post Post

		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}

		res = append(res, &post)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}
