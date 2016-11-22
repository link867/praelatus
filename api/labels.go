package api

import (
	"encoding/json"
	"net/http"
)

func InitLabelRoutes() {
	BaseRoutes.Tickets.Handle("/", mw.Default(GetAllLabels)).Methods("GET")
	BaseRoutes.Tickets.Handle("/", mw.Default(CreateLabel)).Methods("POST")
	BaseRoutes.Tickets.Handle("/{id}", mw.Default(DeleteLabel)).Methods("DELETE")
	BaseRoutes.Tickets.Handle("/{id}", mw.Default(UpdateLabel)).Methods("PUT")
	BaseRoutes.Tickets.Handle("/search", mw.Default(SearchLabels)).Methods("GET")
}

// GetAllLabels will return a JSON array of all labels from the store.
func GetAllLabels(w http.ResponseWriter, r *http.Request) {
	labels, err := Store.Labels().GetAll()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	jsn, err := json.Marshal(labels)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsn)
}
