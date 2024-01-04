package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("it gets Alice's score", func(t *testing.T) {
		request := newGetScoreRequest("Alice")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("it gets Bob's score", func(t *testing.T) {
		request := newGetScoreRequest("Bob")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest("http.MethodGet", fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Response body is wrong, got: %q, but wanted: %q", got, want)
	}
}
