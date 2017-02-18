package cli

import (
	"fmt"

	"github.com/praelatus/praelatus/config"
	"github.com/praelatus/praelatus/store"
	"github.com/urfave/cli"
)

func seedDB(c *cli.Context) error {
	s := config.Store()
	return store.SeedAll(s)
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
