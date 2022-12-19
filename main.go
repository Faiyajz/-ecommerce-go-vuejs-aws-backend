package main

import (
	"backend/internal/server"
	"log"
)

func main() {
	myServer, err := server.New(server.Config{
		Port: 9000,
	})

	if err != nil {
		log.Fatalf("impossible to create the server: %s", err)
	}

	err = myServer.Run()
	if err != nil {
		log.Fatalf("impossible to start the server: %s", err)
	}
}
