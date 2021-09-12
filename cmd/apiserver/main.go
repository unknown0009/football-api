package main

import (
	"log"

	"github.com/fidesy/football-api/internal/apiserver"
)

func main() {
	config := &apiserver.Config{
		BindAddr: ":80",
		DatabaseURL: YOUR_DATABASE_URL,
	}

	s := apiserver.New()
	if err := s.Start(config); err != nil {
		log.Fatal(err)
	}
}
