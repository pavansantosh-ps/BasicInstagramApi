package main

// func TestCreatePostData(t *testing.T) {

// 	t.Run("Inserts a single document", func(t *testing.T) {
// 		var user = []byte(`{
// 			"postid": "51",
// 			"userid": "23",
// 			"caption": "pavan",
// 			"imageurl": "/15/3456/home"
// 		}`)
// 		request, _ := http.NewRequest(http.MethodPost, "/posts/", bytes.NewBuffer(user))
// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		apis.Posts(response, request)
// 		got := response.Body.String()
// 		want := `{"postid": "51"}`

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// func TestGetPostData(t *testing.T) {
// 	t.Run("returns user data by id", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/posts/", nil)
// 		response := httptest.NewRecorder()
// 		apis.Posts(response, request)
// 		got := response.Body.String()
// 		want := `{"postid": "51","userid": "23","caption": "pavan","imageurl": "/15/3456/home","timestamp": "2021-10-09T08:53:44.519Z"}`
// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }
