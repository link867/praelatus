package config

import "os"

// GetDbURL will return the environment variable PRAELATUS_DB if set, otherwise
// return the default development database url.
func GetDbURL() string {
	url := os.Getenv("PRAELATUS_DB")
	if url == "" {
		return "postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable"
	}

	return url
}

// IsDevEnv will return a boolean indicating whether the app is runnning in dev
// mode or not
func IsDevEnv() bool {
	dev := os.Getenv("PRAELATUS_DEV_MODE")
	if dev == "" {
		return false
	}

	return true
}
