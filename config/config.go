package config

import (
	"fmt"
	"io"
	"log/syslog"
	"os"
	"strings"

	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

// GetDbURL will return the environment variable PRAELATUS_DB if set, otherwise
// return the default development database url.
func GetDbURL() string {
	url := os.Getenv("PRAELATUS_DB")
	if url == "" {
		return "postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable"
	}

	return url
}

// Dev will return a boolean indicating whether the app is runnning in dev
// mode or not
func Dev() bool {
	dev := os.Getenv("PRAELATUS_DEV_MODE")
	if dev == "" {
		return false
	}

	return true
}

// Port will return the port / interfaces for the api to listen on based on the
// configuration
func Port() string {
	port := os.Getenv("PRAELATUS_PORT")
	if port == "" {
		return ":8080"
	}

	return port
}

// Store will return the correct data store based on the configuration of the
// instance
func Store() store.Store {
	return pg.New(GetDbURL())
}

func logging() []string {
	loggers := os.Getenv("PRAELATUS_LOGGERS")

	if loggers == "" {
		return []string{"stdout"}
	}

	return strings.Split(loggers, ",")
}

// LogWriter will return an io.writer that is created based on environment
// configuration
func LogWriter() io.Writer {
	if Dev() || os.Getenv("PRAELATUS_LOGGERS") == "" {
		return os.Stdout
	}

	var writers []io.Writer

	for _, log := range logging() {
		switch log {
		case "stdout":
			writers = append(writers, os.Stdout)
		case "syslog":
			sl, err := syslog.New(6, "PRA")
			if err != nil {
				fmt.Println(err)
				continue
			}

			writers = append(writers, sl)
		default:
			var f *os.File
			var e error

			if _, err := os.Stat(log); err == nil {
				f, e = os.Open(log)
			} else {
				f, e = os.Create(log)
			}

			if e != nil {
				fmt.Println(e)
				continue
			}

			writers = append(writers, f)
		}
	}

	return io.MultiWriter(writers...)
}
