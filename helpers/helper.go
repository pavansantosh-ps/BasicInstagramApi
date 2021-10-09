package helper

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ConnectDB() *mongo.Database {
	//creates a connection to mongodb hosted in mongodb://localhost:27017 and returns the client
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
	return db
}
