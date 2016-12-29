package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/mw"
)

func initLabelRoutes() {
	Router.Handle("/labels", mw.Default(GetAllLabels)).Methods("GET")
	Router.Handle("/labels", mw.Default(CreateLabel)).Methods("POST")

	Router.Handle("/labels/{id}", mw.Default(GetLabel)).Methods("GET")
	Router.Handle("/labels/{id}", mw.Default(DeleteLabel)).Methods("DELETE")
	Router.Handle("/labels/{id}", mw.Default(UpdateLabel)).Methods("PUT")
}

// GetAllLabels will return a JSON array of all labels from the store.
func GetAllLabels(w http.ResponseWriter, r *http.Request) {
	labels, err := Store.Labels().GetAll()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, labels)
}

// GetLabel will return a JSON representation of a label
func GetLabel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/labels/"):]

	lbl := &models.Label{}

	i, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError("invalid id"))
		log.Println(err)
		return
	}

	lbl.ID = int64(i)

	err = Store.Labels().Get(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		return
	}

	sendJSON(w, lbl)
}

// CreateLabel creates a label in the db and return a JSON object of
func CreateLabel(w http.ResponseWriter, r *http.Request) {
	var lbl models.Label

	u := mw.GetUser(r.Context())
	if u == nil {
		w.WriteHeader(403)
		w.Write(apiError("you must be logged in to create a label"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lbl)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError("malformed json"))
		log.Println(err)
		return
	}

	err = Store.Labels().New(&lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, lbl)
}

// UpdateLabel updates the label in the db and returns a message indicating
// success or failure.
func UpdateLabel(w http.ResponseWriter, r *http.Request) {
	var lbl models.Label

	u := mw.GetUser(r.Context())
	if u == nil {
		w.WriteHeader(403)
		w.Write(apiError("you must be logged in to create a label"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lbl)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError("malformed json"))
		log.Println(err)
		return
	}

	err = Store.Labels().Save(lbl)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, lbl)
}

// DeleteLabel deletes labels from the db and returns a repsonse indicating
// success of failure.
func DeleteLabel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError("invalid id"))
		log.Println(err)
		return
	}

	err = Store.Labels().Remove(models.Label{ID: int64(i)})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte{})
}
