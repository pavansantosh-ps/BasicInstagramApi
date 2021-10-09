package apis

import (
	"context"
	"encoding/json"
	"fmt"
	helper "internship/instragram/helpers"
	"internship/instragram/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB := helper.ConnectDB()
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
