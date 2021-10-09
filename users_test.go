package main

import (
	"internship/instragram/apis"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/16", nil)
		response := httptest.NewRecorder()
		apis.Users(response, request)
		got := response.Body.String()
		want := ""

		// if got != want {
		// 	t.Errorf("got %q, want %q", got, want)
		// }
	})
}
