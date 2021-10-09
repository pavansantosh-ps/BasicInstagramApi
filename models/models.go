package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type Post struct {
	Postid    primitive.ObjectID `json:"id" bson:"_id"`
	Userid    string             `json:"userid" bson:"userid"`
	Caption   string             `json:"caption" bson:"caption"`
	Imageurl  string             `json:"imageurl" bson:"imageurl"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
