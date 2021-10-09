package models

import "time"

type User struct {
	Id       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Post struct {
	Postid    string    `json:"postid" bson:"postid"`
	Userid    string    `json:"userid" bson:"userid"`
	Caption   string    `json:"caption" bson:"caption"`
	Imageurl  string    `json:"imageurl" bson:"imageurl"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
