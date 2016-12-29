package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

// Router is used to store all the various resource routes.
var Router *mux.Router

// Store is the global store used in our HTTP handlers.
var Store store.Store

// Cache is the global cache object used in our HTTP handlers.
var Cache *store.Cache

// Run will start running the api on the given port
func Run(port string) {
	Store = pg.New(os.Getenv("PRAELATUS_DB"))

	Router = mux.NewRouter()

	initUserRoutes()
	initProjectRoutes()
	initTicketRoutes()
	initFieldRoutes()
	initTypeRoutes()
	initTeamRoutes()
	initWorkflowRoutes()
	initLabelRoutes()

	http.ListenAndServe(port, Router)
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
