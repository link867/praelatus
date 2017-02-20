// Package config defines the config struct and provides utility methods for
// querying it. Additionally it handles loading the config.json if present
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/praelatus/praelatus/store"
	"github.com/praelatus/praelatus/store/pg"
	"github.com/praelatus/praelatus/store/session/bolt"
)

// Config holds much of the configuration for praelatus, if reading from the
// configuration you should use the helper methods in this package as they do
// some prequisite processing and return appropriate types.
type Config struct {
	DBURL        string
	Port         string
	DataDir      string
	LogLocations string
	DevMode      bool
}

func (c Config) String() string {
	return fmt.Sprintf(
		"{\n\tDBURL: %s\n\tPort: %s\n\tDataDir: %s\n\tLogLocations: %s\n\tDevMode: %v\n}",
		c.DBURL, c.Port, c.DataDir, c.LogLocations, c.DevMode)
}

// Cfg is the global config variable used in the helper methods of this package
var Cfg Config

// TODO make this cross platform
func init() {
	Cfg.DBURL = os.Getenv("PRAELATUS_DB")
	if Cfg.DBURL == "" {
		Cfg.DBURL = "postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable"
	}

	dev := os.Getenv("PRAELATUS_DEV_MODE")
	if dev == "" {
		Cfg.DevMode = false
	} else {
		Cfg.DevMode = true
	}

	Cfg.Port = os.Getenv("PRAELATUS_PORT")
	if Cfg.Port == "" {
		Cfg.Port = ":8080"
	}

	Cfg.LogLocations = os.Getenv("PRAELATUS_LOGGERS")
	if Cfg.LogLocations == "" {
		Cfg.LogLocations = "stdout"
	}

	CfgFile := path.Join(os.Getenv("HOME"), ".praelatus")
	f, err := os.Open(CfgFile)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	defer f.Close()

	if os.IsNotExist(err) {
		f, _ = os.Create(CfgFile)

		byt, err := json.Marshal(Cfg)
		if err != nil {
			panic(err)
		}

		f.Write(byt)
		return
	}

	jsn, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var c Config

	err = json.Unmarshal(jsn, &c)
	if err != nil {
		fmt.Println("Error unmarshaling config file defaulting to environment variable configuration")
		fmt.Println(err)
		return
	}

	Cfg = c
}

// GetDbURL will return the environment variable PRAELATUS_DB if set, otherwise
// return the default development database url.
func GetDbURL() string {
	return Cfg.DBURL
}

// Dev will return a boolean indicating whether the app is runnning in dev
// mode or not
func Dev() bool {
	return Cfg.DevMode
}

// Port will return the port / interfaces for the api to listen on based on the
// configuration
func Port() string {
	return Cfg.Port
}

// Store will return the correct data store based on the configuration of the
// instance
func Store() store.Store {
	return pg.New(GetDbURL())
}

// SessionStore will return a session store with a default location
func SessionStore() store.SessionStore {
	return bolt.New("sessions.db")
}

// GetRedis will get the redis cache location
func GetRedis() string {
	return ""
}
