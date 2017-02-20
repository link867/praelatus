// Praelatus is an Open Source bug tracking and ticketing system. The backend
// API is written in Go and the frontend is a React.js app. You are viewing the
// Godoc for the backend.
package main

import (
	"os"

	"github.com/praelatus/praelatus/cli"
)

func main() {
	cli.Run(os.Args)
}
