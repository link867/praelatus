package pg

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg/migrations"
)

type rowScanner interface {
	Scan(dest ...interface{}) error
}

// Store implements the store.Store and store.SQLStore interface for a postgres DB.
type Store struct {
	db        *sql.DB
	replicas  []sql.DB
	users     *UserStore
	projects  *ProjectStore
	fields    *FieldStore
	workflows *WorkflowStore
	tickets   *TicketStore
	types     *TypeStore
	labels    *LabelStore
	statuses  *StatusStore
	teams     *TeamStore
}

// New connects to the postgres database provided and returns a store
// that's connected.
func New(conn string, replicas ...string) store.Store {

	d, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panicln("Error connection:", err)
	}

	s := &Store{
		db:        d,
		replicas:  []sql.DB{},
		users:     &UserStore{d},
		projects:  &ProjectStore{d},
		fields:    &FieldStore{d},
		tickets:   &TicketStore{d},
		labels:    &LabelStore{d},
		workflows: &WorkflowStore{d},
		types:     &TypeStore{d},
		statuses:  &StatusStore{d},
		teams:     &TeamStore{d},
	}

	err = migrations.RunMigrations(s.db)
	if err != nil {
		log.Panicln("Error migrating:", err)
	}

	return s
}

// Users returns the underlying UserStore for a postgres DB
func (pg *Store) Users() store.UserStore {
	return pg.users
}

// Teams returns the underlying TeamStore for a postgres DB
func (pg *Store) Teams() store.TeamStore {
	return pg.teams
}

// Fields returns the underlying FieldStore for a postgres DB
func (pg *Store) Fields() store.FieldStore {
	return pg.fields
}

// Tickets returns the underlying TicketStore for a postgres DB
func (pg *Store) Tickets() store.TicketStore {
	return pg.tickets
}

// Types returns the underlying TypeStore for a postgres DB
func (pg *Store) Types() store.TypeStore {
	return pg.types
}

// Projects returns the underlying ProjectStore for a postgres DB
func (pg *Store) Projects() store.ProjectStore {
	return pg.projects
}

// Statuses returns the underlying StatusStore for a postgres DB
func (pg *Store) Statuses() store.StatusStore {
	return pg.statuses
}

// Workflows returns the underlying WorkflowStore for a postgres DB
func (pg *Store) Workflows() store.WorkflowStore {
	return pg.workflows
}

// Labels returns the underlying LabelStore for a postgres DB
func (pg *Store) Labels() store.LabelStore {
	return pg.labels
}

// Conn implementes store.SQLStore for postgres db
func (pg *Store) Conn() *sql.DB {
	return pg.db
}

// toPqErr converts an error to a pq.Error so we can access more info about what
// happened.
func toPqErr(e error) *pq.Error {
	if err, ok := e.(*pq.Error); ok {
		return err
	}

	return nil
}

func handlePqErr(e error) error {
	pqe := toPqErr(e)
	if pqe == nil {
		return e
	}

	log.Printf("pq error [%v] %s\n", pqe.Code, pqe.Message)

	// fmt.Println("PQ ERROR CODE:", pqe.Code)
	if pqe.Code == "23505" {
		return store.ErrDuplicateEntry
	}

	return e
}
