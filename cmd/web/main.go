package main

import (
	"linksh/internal/webserver"
	"log"
)

func main() {
	server := webserver.Server{Port: ":8080"}
	err := server.Run()
	if err != nil {
		log.Fatalf("cannot run web server on port :8080, reason: %s", err.Error())
	}
}
