package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/praelatus/backend/mw"
	"github.com/praelatus/backend/store"
	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
)

// Store is the global store used in our HTTP handlers.
var Store store.Store

// Cache is the global cache object used in our HTTP handlers.
var Cache *store.Cache

func index(rtr chi.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsnStr := docgen.JSONRoutesDoc(rtr)
		w.Write([]byte(jsnStr))
	})
}

// New will start running the api on the given port
func New(store store.Store) chi.Router {
	Store = store

	router := chi.NewRouter()

	router.Use(mw.Default...)
	// router.Use(middleware.Recoverer)

	api := chi.NewRouter()

	api.Mount("/routes", index(api))
	api.Mount("/fields", fieldRouter())
	api.Mount("/labels", labelRouter())
	api.Mount("/projects", projectRouter())
	api.Mount("/teams", teamRouter())
	api.Mount("/tickets", ticketRouter())
	api.Mount("/types", typeRouter())
	api.Mount("/users", userRouter())
	api.Mount("/workflows", workflowRouter())

	router.Mount("/api", api)
	router.Mount("/api/v1", api)

	// Left here for debugging purposes
	// docgen.PrintRoutes(router)

	return router
}

// Message is a general purpose json struct used primarily for error responses.
type Message struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func apiError(msg string, fields ...string) []byte {
	e := Message{
		Message: msg,
	}

	if fields != nil {
		e.Field = strings.Join(fields, ",")
	}

	byt, _ := json.Marshal(e)
	return byt
}

func sendJSON(w http.ResponseWriter, v interface{}) {
	resp, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError("Failed to marshal database response to JSON."))
		log.Println(err)
		return
	}

	w.Write(resp)
}
