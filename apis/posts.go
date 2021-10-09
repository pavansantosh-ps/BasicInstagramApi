package apis

import (
	"context"
	"encoding/json"
	"fmt"
	helper "internship/instragram/helpers"
	"internship/instragram/models"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB := helper.ConnectDB()
	switch r.Method {
	case "GET":
		postID := r.URL.Path[len("/posts/"):]
		postsCollection := DB.Collection("posts")
		objID, _ := primitive.ObjectIDFromHex(postID)
		result := postsCollection.FindOne(context.TODO(), bson.M{"_id": objID})
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
		p.Postid = primitive.NewObjectID()
		postsCollection := DB.Collection("posts")
		insertResult, err := postsCollection.InsertOne(context.TODO(), p)
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"postid": "%s"}`, p.Postid.Hex())))
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
		skip := int64(0)
		limit := int64(5)
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}
		size, err := strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil {
			size = 5
		}
		fmt.Print(page)
		fmt.Print(size)
		if page <= 0 {
			page = 1
		}
		if size <= 0 {
			size = 5
		}
		skip = int64((page - 1) * size)
		limit = int64(size)

		opts := options.FindOptions{}
		opts.Skip = &skip
		opts.Limit = &limit
		curser, err := postsCollection.Find(context.TODO(), bson.M{"userid": userId}, &opts)

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
