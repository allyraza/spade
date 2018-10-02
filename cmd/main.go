package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/allyraza/spade"
)

func main() {
	config := &spade.Config{}

	flag.StringVar(&config.Address, "address", ":2000", "address to listen on")
	flag.Parse()

	app := spade.New(config)

	log.Printf("listening on %v\n", config.Address)

	err := http.ListenAndServe(config.Address, app.Mux)
	if err != nil {
		log.Fatalf("Spade HTTP: %v\n", err)
	}
}
