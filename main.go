package main

import (
	"fmt"
	"log"
	"net/http"
)

// PlayerStore records player scores
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer returns a user score
type PlayerServer struct {
	store PlayerStore
}

// ServerHTTP servers player scores
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPlayerScore(player))
}

// GetPlayerScore returns the score for the specified player
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

func main() {
	server := &PlayerServer{}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
