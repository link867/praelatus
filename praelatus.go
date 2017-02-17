package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/praelatus/praelatus/api"
	"github.com/praelatus/praelatus/cli"
	"github.com/praelatus/praelatus/config"
	"github.com/praelatus/praelatus/models"
	"github.com/praelatus/praelatus/store"
	"github.com/tylerb/graceful"
)

func main() {
	cli.Run(os.Args)
}

func adminCreateUser(c *cli.Context) error {
	username := c.String("username")
	password := c.String("password")
	fullName := c.String("fullName")
	email := c.String("email")
	admin := c.Bool("admin")

	if username == "" {
		return cli.NewExitError("missing required --username flag", 1)
	}

	if password == "" {
		return cli.NewExitError("missing required --password flag", 1)
	}

	if fullName == "" {
		return cli.NewExitError("missing required --fullName flag", 1)
	}

	if email == "" {
		return cli.NewExitError("missing required --email flag", 1)
	}

	u, err := models.NewUser(username, password, fullName, email, admin)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	s := config.Store()
	if err := s.Users().New(u); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}

func seedDB(c *cli.Context) error {
	s := config.Store()
	return store.SeedAll(s)
}

func showConfig(c *cli.Context) error {
	fmt.Println(config.Cfg)
	return nil
}

func testDB(c *cli.Context) error {
	// these methods panic if connection fails
	config.Store()
	config.SessionStore()
	return nil
}

func cleanDB(c *cli.Context) error {
	s := config.Store()
	sql, ok := s.(store.SQLStore)
	if !ok {
		fmt.Println("Configured database is not sql nothing to do.")
		return nil
	}

	err := sql.Drop()
	if err != nil {
		return err
	}

	fmt.Println("Database cleaned.")
	return nil
}

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

func runServer(c *cli.Context) error {
	log.SetOutput(config.LogWriter())

	log.Println("Starting Praelatus...")
	log.Println("Initializing database...")

	s := config.Store()
	ss := config.SessionStore()

	if sql, ok := s.(store.SQLStore); ok {
		log.Println("Migrating database...")
		err := sql.Migrate()
		if err != nil {
			log.Println("Error migrating database:", err)
			os.Exit(1)
		}
	}

	log.Println("Prepping API")
	var r http.Handler = api.New(s, ss)
	if config.Dev() || c.Bool("devmode") {
		log.Println("Running in dev mode, disabling cors...")
		r = disableCors(r)
	}

	log.Println("Ready to serve requests!")
	return graceful.RunWithErr(config.Port(), time.Minute, r)
}
