package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, fmt.Sprintf("%d", p.store.GetPlayerScore(player)))
}
