package main

import (
	"log"
	"net/http"
)

func main() {
	store, err := NewBoltPlayerStore("players.db")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
