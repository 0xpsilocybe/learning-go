package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type VolatilePlayerStore struct {
	scores map[string]int
}

func (v *VolatilePlayerStore) GetPlayerScore(name string) int {
	return v.scores[name]
}

func (v *VolatilePlayerStore) RecordWin(name string) {
	v.scores[name]++
}

func NewVolatilePlayerStore() *VolatilePlayerStore {
	return &VolatilePlayerStore{map[string]int{}}
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := p.store.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, name string) {
	p.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
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
