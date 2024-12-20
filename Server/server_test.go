package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestGETPlayers(t *testing.T) {
// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		got := response.Body.String()
// 		want := "20"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// 	t.Run("returns Floyd's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		got := response.Body.String()
// 		want := "10"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}