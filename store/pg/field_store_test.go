package pg_test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/praelatus/backend/config"
	"github.com/praelatus/backend/models"
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

func TestFieldGet(t *testing.T) {
	f := &models.Field{ID: 1}
	e := s.Fields().Get(f)
	failIfErr("Field Get", t, e)

	if f.Name == "" {
		t.Error("Expected a name got %s\n", f.Name)
	}
}

func TestFieldGetAll(t *testing.T) {
	f, e := s.Fields().GetAll()
	failIfErr("Field Get All", t, e)

	if f == nil || len(f) == 0 {
		t.Error("Expected multiple fields and got nil instead.")
	}
}

func TestFieldGetByProject(t *testing.T) {
	p := &models.Project{ID: 1}

	f, e := s.Fields().GetByProject(p)
	failIfErr("Field Get By Project", t, e)

	if f == nil || len(f) == 0 {
		t.Error("Expected multiple fields and got nil instead.")
	}
}

func TestFieldSave(t *testing.T) {
	field := &models.Field{
		ID:       1,
		Name:     "Story Points",
		DataType: "INT",
	}

	e := s.Fields().Save(field)
	failIfErr("Field Save", t, e)

	f, e := s.Fields().Get(1)
	failIfErr("Field Save", t, e)

	if f.Name != "Story Points" {
		t.Errorf("Expected Story Points got: %s\n", f.Name)
	}

	if f.DataType != "INT" {
		t.Errorf("Expected INT got: %s\n", f.DataType)
	}
}

func TestFieldRemove(t *testing.T) {
	field := &models.Field{
		ID:   2,
		Name: "TestField2",
	}

	e := s.Fields().Remove(field)
	failIfErr("Field Remove", t, e)
}
