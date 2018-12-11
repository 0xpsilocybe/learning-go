package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bartek/learning-go/bserv"
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
	poker.NewCLI(db, os.Stdin).PlayPoker()
}
