package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/praelatus/backend/api"
	"github.com/praelatus/backend/config"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
	"github.com/tylerb/graceful"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "praelatus"
	app.Usage = "Praelatus, an Open Source bug tracker / ticketing system"
	app.Version = "0.0.1"
	app.Action = runServer
	app.Authors = []cli.Author{
		{
			Name:  "Mathew Robinson",
			Email: "chasinglogic@gmail.com",
		},
		{
			Name:  "Mark Chandler",
			Email: "mark.allen.chandler@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "seed",
			Usage:  "seed the database with test data",
			Action: seed,
		},
		{
			Name:   "serve",
			Usage:  "start running the REST api",
			Action: runServer,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "devmode",
					Usage: "runs server in devmode which changes some security behavior to ease development",
				},
			},
		},
		{
			Name:   "config",
			Action: showConfig,
			Usage:  "view the configuration for this instance, useful for debugging",
		},
		{
			Name:  "admin",
			Usage: "various admin functions for the instance",
			Subcommands: []cli.Command{
				{
					Name:   "createUser",
					Usage:  "create a user, useful for creating admin accounts",
					Action: adminCreateUser,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "username",
						},
						cli.StringFlag{
							Name: "password",
						},
						cli.StringFlag{
							Name: "fullName",
						},
						cli.StringFlag{
							Name: "email",
						},
						cli.BoolFlag{
							Name:  "admin",
							Usage: "when this flag is given user will be created as an system admin",
						},
					},
				},
			},
		},
		{
			Name:   "testdb",
			Usage:  "will test the connections to the databases",
			Action: testDB,
		},
	}

	app.Run(os.Args)
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

func seed(c *cli.Context) error {
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

	log.Println("Prepping API")
	var r http.Handler = api.New(s, ss)
	if config.Dev() || c.Bool("devmode") {
		log.Println("Running in dev mode, disabling cors...")
		r = disableCors(r)
	}

	log.Println("Ready to serve requests!")
	return graceful.RunWithErr(config.Port(), time.Minute, r)
}
