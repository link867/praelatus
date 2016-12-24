package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/mw"
)

func initLabelRoutes() {
	Router.Handle("/labels", mw.Default(GetAllLabels)).Methods("GET")
	Router.Handle("/labels", mw.Default(CreateLabel)).Methods("POST")
	Router.Handle("/labels/{idOrName}", mw.Default(GetLabel)).Methods("GET")
	Router.Handle("/labels/{idOrName}", mw.Default(DeleteLabel)).Methods("DELETE")
	Router.Handle("/labels/{idOrName}", mw.Default(UpdateLabel)).Methods("PUT")
	Router.Handle("/labels/search", mw.Default(SearchLabels)).Methods("GET")
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

// GetLabel will return a JSON representation of a model.
func GetLabel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lbl := &models.Label{}

	i, err := strconv.Atoi(vars["idOrName"])
	if err != nil {
		lbl.Name = vars["idOrName"]
	}

	lbl.ID = int64(i)

	err = Store.Labels().Get(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	jsn, err := json.Marshal(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsn)
}

// CreateLabel creates a label in the db and return a JSON object of
func CreateLabel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lbl := &models.Label{}

	i, err := strconv.Atoi(vars["idOrName"])
	if err != nil {
		lbl.Name = vars["idOrName"]
	}

	lbl.ID = int64(i)

	err = Store.Labels().New(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(lbl.String()))
}

// UpdateLabel updates the label in the db and returns a message indicating
// success or failure.
func UpdateLabel(w http.ResponseWriter, r *http.Request) {
	lbl := models.Label{}

	err := Store.Labels().Save(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Label successfully updated"))

}

// DeleteLabel deletes labels from the db and returns a repsonse indicating
// success of failure.
func DeleteLabel(w http.ResponseWriter, r *http.Request) {
	lbl := models.Label{}

	err := Store.Labels().Remove(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Label successfully deleted"))
}

// SearchLabels
func SearchLabels(w http.ResponseWriter, r *http.Request) {

}
