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

func TestCreatePostData(t *testing.T) {
	//Testing POST method by creating a post by making an http request
	t.Run("Inserts a single post", func(t *testing.T) {
		var user = []byte(`{
			"userid": "61616cfe46c8143bf86d96f3",
			"caption": "pavan",
			"imageurl": "https:/google.com"
		}`)
		request, _ := http.NewRequest(http.MethodPost, "/posts/", bytes.NewBuffer(user))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		apis.Posts(response, request)

		if response.Code >= 400 {
			t.Errorf("resp %q", response.Code)
		}
	})
}

func TestGetPostData(t *testing.T) {
	//Testing Get method by creating a post by making http request and checking if the post exists in the posts document
	t.Run("Creates a post and checks if the post exists", func(t *testing.T) {
		var post = []byte(`{
			"userid": "61616cfe46c8143bf86d96f3",
			"caption": "pavan",
			"imageurl": "https:/google.com"
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/posts/", bytes.NewBuffer(post))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		apis.Posts(res, req)
		postResp := res.Body.String()
		var respMap map[string]string
		err := json.Unmarshal([]byte(postResp), &respMap)
		if err != nil {
			t.Errorf("error parsing the resp %q", postResp)
		}
		postID := respMap["postid"]
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/posts/%s", postID), nil)
		response := httptest.NewRecorder()
		apis.Posts(response, request)
		got := response.Body.String()
		var gotMap map[string]string
		err = json.Unmarshal([]byte(got), &gotMap)
		if err != nil {
			t.Errorf("error parsing the resp %q", postResp)
		}
		if gotMap["id"] != postID || gotMap["userid"] != "61616cfe46c8143bf86d96f3" || gotMap["caption"] != "pavan" || gotMap["imageurl"] != "https:/google.com" {
			t.Errorf("got name:'%s' email:'%s', want name: santosh email: santosh@gmail.com", gotMap["name"], gotMap["email"])
		}
	})
}
