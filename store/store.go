package store

import (
	"database/sql"
	"errors"

	"github.com/praelatus/backend/models"
)

var (
	// ErrDuplicateEntry is returned when a unique constraint is violated.
	ErrDuplicateEntry = errors.New("duplicate entry attempted")
	// ErrNotFound is returned when an invalid resource is given or searched
	// for
	ErrNotFound = errors.New("no such resource")
)

// Store is an interface for storing and retrieving models.
type Store interface {
	Users() UserStore
	Teams() TeamStore
	Labels() LabelStore
	Fields() FieldStore
	Tickets() TicketStore
	Types() TypeStore
	Projects() ProjectStore
	Statuses() StatusStore
	Workflows() WorkflowStore
}

// SQLStore is an interface for a sql store so we can request direct
// access to the database.
type SQLStore interface {
	Conn() *sql.DB
}

// Cache is an abstraction over using Redis or any other caching system.
type Cache interface {
	Get(string) interface{}
	Set(string, interface{}) error
}

// FieldStore contains methods for storing and retrieving Fields and FieldValues
type FieldStore interface {
	Get(*models.Field) error
	GetAll() ([]models.Field, error)

	GetByProject(models.Project) ([]models.Field, error)
	AddToProject(project models.Project, field *models.Field,
		ticketTypes ...models.TicketType) error

	New(*models.Field) error
	Save(models.Field) error
	Remove(models.Field) error
}

// UserStore contains methods for storing and retrieving Users
type UserStore interface {
	Get(*models.User) error
	GetAll() ([]models.User, error)

	New(*models.User) error
	Save(models.User) error
	Remove(models.User) error
}

// ProjectStore contains methods for storing and retrieving Projects
type ProjectStore interface {
	Get(*models.Project) error
	GetAll() ([]models.Project, error)

	New(*models.Project) error
	Save(models.Project) error
	Remove(models.Project) error
}

// TypeStore is used to save and retrieve Ticket Types
type TypeStore interface {
	Get(*models.TicketType) error
	GetAll() ([]models.TicketType, error)

	New(*models.TicketType) error
	Save(models.TicketType) error
	Remove(models.TicketType) error
}

// TicketStore contains methods for storing and retrieving Tickets
type TicketStore interface {
	Get(models.Project, *models.Ticket) error
	GetAll() ([]models.Ticket, error)
	GetAllByProject(models.Project) ([]models.Ticket, error)

	GetComments(models.Ticket) ([]models.Comment, error)
	NewComment(models.Ticket, *models.Comment) error
	SaveComment(models.Comment) error
	RemoveComment(models.Comment) error

	NextTicketKey(models.Project) string

	New(models.Project, *models.Ticket) error
	Save(models.Ticket) error
	Remove(models.Ticket) error
}

// TeamStore contains methods for storing and retrieving Teams
type TeamStore interface {
	Get(*models.Team) error
	GetAll() ([]models.Team, error)
	GetForUser(models.User) ([]models.Team, error)

	AddMembers(models.Team, ...models.User) error
	GetMembers(*models.Team) error

	New(*models.Team) error
	Save(models.Team) error
	Remove(models.Team) error
}

// StatusStore contains methods for storing and retrieving Statuses
type StatusStore interface {
	Get(*models.Status) error
	GetAll() ([]models.Status, error)

	New(*models.Status) error
	Save(models.Status) error
	Remove(models.Status) error
}

// WorkflowStore contains methods for storing and retrieving Workflows
type WorkflowStore interface {
	Get(models.Project, *models.Workflow) error
	GetAll() ([]models.Workflow, error)
	GetByProject(models.Project) ([]models.Workflow, error)

	New(models.Project, *models.Workflow) error
	Save(models.Workflow) error
	Remove(models.Workflow) error
}

// LabelStore contains methods for storing and retrieving Labels
type LabelStore interface {
	Get(*models.Label) error
	GetAll() ([]models.Label, error)

	New(*models.Label) error
	Save(models.Label) error
	Remove(models.Label) error
}
