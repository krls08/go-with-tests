package main

import (
	"log"
	"net/http"

	poker "github.com/krls08/go-with-tests/23_APP/27_Command_line_package_structure"
)

const dbFileName = "game.db.json"

func main() {
	//server := &PlayerServer{store: NewInMemoryPlayerStore()}
	//server := NewPlayerServer(NewInMemoryPlayerStore())

	anStore, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	closeDB()

	server := poker.NewPlayerServer(anStore)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
