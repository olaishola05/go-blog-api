package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type SocialLinks struct {
	Twitter   string `bson:"twitter,omitempty" json:"twitter,omitempty"`
	LinkedIn  string `bson:"linkedin,omitempty" json:"linkedin,omitempty"`
	GitHub    string `bson:"github,omitempty" json:"github,omitempty"`
	Portfolio string `bson:"portfolio,omitempty" json:"portfolio,omitempty"`
}

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	Role      string        `bson:"role" json:"role"`
	Bio       string        `bson:"bio,omitempty" json:"bio,omitempty"`
	Socials   SocialLinks   `bson:"socials,omitempty" json:"socials,omitempty"`
	CreatedAt bson.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt bson.DateTime `bson:"updated_at" json:"updated_at"`
}

var mgo MongoInstance

func (user *User) CreateUser() (*User, error) {
	return user, nil
}

func GetAllUsers(ctx context.Context) ([]User, error) {
	mgo := GetDb()
	var users []User
	query := bson.D{{}}
	cursor, err := mgo.Db.Collection("user").Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id string) {

}

func UpdateUser(id string) {

}

func DeleteUser(id string) {

}
