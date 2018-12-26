package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0xpsilocybe/learning-go/bserv"
)

const connection = "../../players.db"

func main() {
	db, err := poker.NewBoltPlayerStore(connection)
	defer db.Close()
	if err != nil {
		log.Fatalf("There was a problem creating a database %v", err)
	}
	fmt.Println("Let's play poker!")
	fmt.Println("Type ''{name} wins' to record a win")
	game := poker.NewGame(
		poker.BlindAlerterFunc(poker.StdOutAlerter),
		db,
	)
	poker.NewCLI(
		os.Stdin,
		os.Stdout,
		game,
	).PlayPoker()
}
