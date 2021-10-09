package apis

import (
	"context"
	"encoding/json"
	"fmt"
	helper "internship/instragram/helpers"
	"internship/instragram/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB := helper.ConnectDB()
	switch r.Method {
	case "GET":
		//returns the user based on the id provided in the URL
		userId := r.URL.Path[len("/users/"):]
		usersCollection := DB.Collection("users")
		objID, _ := primitive.ObjectIDFromHex(userId)
		result := usersCollection.FindOne(context.TODO(), bson.M{"_id": objID}) //queries the user collection from mongodb by id
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
		//Creates a user document based on the given payload
		decoder := json.NewDecoder(r.Body)
		var u models.User
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}
		u.Password, err = helper.HashPassword(u.Password) //performs hashing of the password
		if err != nil {
			panic(err)
		}
		u.Id = primitive.NewObjectID()
		usersCollection := DB.Collection("users")
		insertResult, err := usersCollection.InsertOne(context.TODO(), u) //inserts the payload into mongodb users collection
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, u.Id.Hex())))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
