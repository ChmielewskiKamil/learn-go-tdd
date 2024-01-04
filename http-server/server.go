package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.String(), "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) (score string) {
	if player == "Bob" {
		score = "10"
	}

	if player == "Alice" {
		score = "20"
	}

	return
}
