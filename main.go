package main

import (
	"internship/instragram/apis"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", apis.Users)
	http.HandleFunc("/posts/users/", apis.GetUserPosts)
	http.HandleFunc("/posts/", apis.Posts)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
