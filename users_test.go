package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"internship/instragram/apis"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatseUserData(t *testing.T) {
	//Testing POST method by creating a user by making an http request
	t.Run("Inserts a single document", func(t *testing.T) {
		var user = []byte(`{
			"name":"santosh",
			"email":"santosh@gmail.com",
			"password":"bchsbchba"}`)
		request, _ := http.NewRequest(http.MethodPost, "/users/", bytes.NewBuffer(user))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		apis.Users(response, request)

		if response.Code >= 400 {
			t.Errorf("resp %q", response.Code)
		}
	})
}

func TestGetUserData(t *testing.T) {
	//Testing Get method by creating a user by making http request and checking if the user exists in the posts document
	t.Run("Creates a user and checks if the user exists", func(t *testing.T) {
		var user = []byte(`{
			"name":"santosh",
			"email":"santosh@gmail.com",
			"password":"bchsbchba"}`)
		req, _ := http.NewRequest(http.MethodPost, "/users/", bytes.NewBuffer(user))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		apis.Users(res, req)
		userResp := res.Body.String()
		var respMap map[string]string
		err := json.Unmarshal([]byte(userResp), &respMap)
		if err != nil {
			t.Errorf("error parsing the resp %q", userResp)
		}
		userID := respMap["id"]
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userID), nil)
		response := httptest.NewRecorder()
		apis.Users(response, request)
		got := response.Body.String()
		var gotMap map[string]string
		err = json.Unmarshal([]byte(got), &gotMap)
		if err != nil {
			t.Errorf("error parsing the resp %q", userResp)
		}
		if gotMap["name"] != "santosh" || gotMap["email"] != "santosh@gmail.com" {
			t.Errorf("got name:'%s' email:'%s', want name: santosh email: santosh@gmail.com", gotMap["name"], gotMap["email"])
		}
	})
}
