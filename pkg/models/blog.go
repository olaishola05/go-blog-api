package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Post struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Content   string        `bson:"content" json:"content"`
	AuthorID  bson.ObjectID `bson:"author_id" json:"author_id"`
	Tags      []string      `bson:"tags,omitempty" json:"tags,omitempty"`
	Published bool          `bson:"published" json:"published"`
	CreatedAt bson.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt bson.DateTime `bson:"updated_at" json:"updated_at"`
}

func GetPosts(ctx context.Context) ([]Post, error) {
	var posts []Post
	return posts, nil
}

func CreatePost(ctx context.Context) (*Post, error) {
	var post Post
	return &post, nil
}

func GetPostById(id string) (*Post, error) {
	var post Post
	return &post, nil
}

func UpdatePostById(id string) (*Post, error) {
	var post Post
	return &post, nil
}

func DeletePostById(id string) (*Post, error) {
	var post Post
	return &post, nil
}
