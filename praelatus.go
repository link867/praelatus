package main

import (
	"os"

	"log"

	"github.com/praelatus/backend/api"
)

func main() {
	log.Println("Starting Praelatus!")
	log.Println("Initializing database...")

	// Will be used for logging
	// mw := io.MultiWriter(os.Stdout)

	log.Println("Ready to serve requests!")
	port := os.Getenv("PRAELATUS_PORT")
	if port == "" {
		port = ":8080"
	}

	api.Run(port)
}
