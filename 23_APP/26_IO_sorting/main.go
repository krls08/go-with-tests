package main

import (
	"log"
	"net/http"
)

func main() {
	//server := &PlayerServer{store: NewInMemoryPlayerStore()}
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
