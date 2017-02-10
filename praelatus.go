package main

import (
	"net/http"

	"log"

	"github.com/praelatus/backend/api"
	"github.com/praelatus/backend/config"
	"github.com/praelatus/backend/store"
)

func main() {
	log.SetOutput(config.LogWriter())

	log.Println("Starting Praelatus!")
	log.Println("Initializing database...")

	s := config.Store()
	ss := config.SessionStore()

	var err error

	if config.Dev() {
		log.Println("Dev environment detected, seeding database with test info...")
		err = store.SeedAll(s)
	} else {
		err = store.SeedDefaults(s)
	}

	if err != nil {
		panic(err)
	}

	log.Println("Prepping API")
	r := api.New(s, ss)

	log.Println("Ready to serve requests!")
	err = http.ListenAndServe(config.Port(), r)
	if err != nil {
		log.Println(err)
	}
}
