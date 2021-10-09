package apis

import (
	"context"
	"encoding/json"
	"fmt"
	helper "internship/instragram/helpers"
	"internship/instragram/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB := helper.ConnectDB()
	switch r.Method {
	case "GET":
		postID := r.URL.Path[len("/posts/"):]
		postsCollection := DB.Collection("posts")
		result := postsCollection.FindOne(context.TODO(), bson.M{"postid": postID})
		post := &models.Post{}
		err := result.Decode(post)
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(post)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var p models.Post
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}
		p.Timestamp = time.Now()
		postsCollection := DB.Collection("posts")
		insertResult, err := postsCollection.InsertOne(context.TODO(), p)
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, p.Postid)))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not posts"}`))
	}
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB := helper.ConnectDB()
	switch r.Method {
	case "GET":
		userId := r.URL.Path[len("/posts/users/"):]
		postsCollection := DB.Collection("posts")
		curser, err := postsCollection.Find(context.TODO(), bson.M{"userid": userId})
		var posts []*models.Post

		if err != nil {
			panic(err)
		}
		for curser.Next(context.TODO()) {
			var result models.Post
			err := curser.Decode(&result)
			if err != nil {
				panic(err)
			}
			posts = append(posts, &result)
		}
		b, err := json.Marshal(posts)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not posts"}`))
	}
}
