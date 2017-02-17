package main

import (
	"net/http"
	"time"

	"log"

	"github.com/praelatus/backend/api"
	"github.com/praelatus/backend/config"
	"github.com/praelatus/backend/store"
	"github.com/tylerb/graceful"
)

// this is only used when running in dev mode to make testing the api easier
func disableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if w.Header().Get("Access-Control-Allow-Origin") == "" {
				w.Header().Add("Access-Control-Allow-Origin", "*")
			}

			next.ServeHTTP(w, r)
		})
}

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
	var r http.Handler = api.New(s, ss)

	if config.Dev() {
		log.Println("Running in dev mode, disabling cors...")
		r = disableCors(r)
	}

	log.Println("Ready to serve requests!")
	err = graceful.RunWithErr(config.Port(), time.Minute, r)
	if err != nil {
		log.Println(err)
	}
}
