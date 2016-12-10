package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

// Routes holds all of the routes for the different API endpoints.
type Routes struct {
	Root     *mux.Router
	Users    *mux.Router
	Projects *mux.Router
	Tickets  *mux.Router
}

// Used in starting the router.
var BaseRoutes *Routes

// Store is the global store used in our HTTP handlers.
var Store store.Store

// Cache is the global cache object used in our HTTP handlers.
var Cache *store.Cache

func Run(port string) {
	Store = pg.New(os.Getenv("PRAELATUS_DB"))

	BaseRoutes = &Routes{}
	BaseRoutes.Root = mux.NewRouter()
	BaseRoutes.Users = BaseRoutes.Root.PathPrefix("/users").Subrouter()
	BaseRoutes.Projects = BaseRoutes.Root.PathPrefix("/projects").Subrouter()
	BaseRoutes.Tickets = BaseRoutes.Root.PathPrefix("/tickets").Subrouter()

	// InitUserRoutes()
	// InitProjectRoutes()
	// InitTicketRoutes()

	http.ListenAndServe(port, BaseRoutes.Root)
}
