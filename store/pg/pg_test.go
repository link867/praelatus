package pg_test

import (
	"fmt"
	"testing"

	"github.com/praelatus/backend/config"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

var s store.Store
var seeded bool

func init() {
	if !seeded {
		fmt.Println("Prepping tests")
		s = pg.New(config.GetDbURL())
		e := store.SeedAll(s)
		if e != nil {
			panic(e)
		}

		seeded = true
	}
}

func failIfErr(testName string, t *testing.T, e error) {
	if e != nil {
		t.Error(testName, " failed with error: ", e)
	}
}
