package main

import (
	"context"
	"encoding/json"
	"fmt"
	helper "internship/instragram/helpers"
	"internship/instragram/models"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("instragram")
	DB = db
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		userId := r.URL.Path[len("/users/"):]
		usersCollection := DB.Collection("users")
		result := usersCollection.FindOne(context.TODO(), bson.M{"id": userId})
		user := &models.User{}
		err := result.Decode(user)
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var u models.User
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}
		u.Password, err = helper.HashPassword(u.Password)
		if err != nil {
			panic(err)
		}
		usersCollection := DB.Collection("users")
		insertResult, err := usersCollection.InsertOne(context.TODO(), u)
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, u.Id)))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		postID := r.URL.Path[len("/posts/"):]
		postsCollection := DB.Collection("posts")
		result := postsCollection.FindOne(context.TODO(), bson.M{"id": postID})
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

func getposts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func main() {
	ConnectDB()
	http.HandleFunc("/users/", users)
	http.HandleFunc("/posts/users/", getposts)
	http.HandleFunc("/posts/", posts)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
