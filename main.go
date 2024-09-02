package main

import (
	"log"

	"github.com/morfo-si/go-microservices/internal/database"
	"github.com/morfo-si/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize the database client: %s", err)
	}
	srv := server.NewEchoServer(db)

	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
