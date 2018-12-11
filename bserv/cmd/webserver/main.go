package main

import (
	"log"
	"net/http"

	"github.com/bartek/learning-go/bserv"
)

const connection = "../../players.db"

func main() {
	db, err := poker.NewBoltPlayerStore(connection)
	if err != nil {
		log.Fatal("There was a problem creating a database %v", err)
	}
	defer db.Close()
	server := poker.NewPlayerServer(db)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
