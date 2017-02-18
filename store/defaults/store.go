// Package defaults provides defaults for stores, for example it will hold
// functions for creating a default Agile Project, Scrum Project, etc. as well
// as defaults for multiple models (issuetypes, labels, etc.) it finally will
// hold a method for getting the default store implementation (postgres)
package defaults

import (
	"github.com/praelatus/praelatus/config"
	"github.com/praelatus/praelatus/store"
	"github.com/praelatus/praelatus/store/pg"
)

// Store returns the default store.Store implementation (postgres)
func Store() store.Store {
	return pg.New(config.GetDbURL())
}
