package main

import (
	"log"

	"github.com/morfo-si/go-microservices/internal/configuration"
	"github.com/morfo-si/go-microservices/internal/database"
	"github.com/morfo-si/go-microservices/internal/server"
)

func main() {
	// Load the configuration
	config := configuration.LoadConfig()

	// Pass the configuration to the database client
	db, err := database.NewDatabaseClient(config)
	if err != nil {
		log.Fatalf("failed to initialize the database client: %s", err)
	}

	srv := server.NewGinServer(db)

	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
