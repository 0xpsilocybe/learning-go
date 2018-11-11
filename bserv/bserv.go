package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type VolatilePlayerStore struct {
}

func (v *VolatilePlayerStore) GetPlayerScore(name string) int {
	return 123
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func GetPlayerScore(name string) int {
	switch name {
	case "Floyd":
		return 10
	case "Pepper":
		return 20
	default:
		return 0
	}
}
