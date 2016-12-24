package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/mw"
)

func initProjectRoutes() {
	Router.Handle("/", mw.Default(GetAllProjects)).Methods("GET")
	Router.Handle("/{pkey}", mw.Default(GetProject)).Methods("GET")
}

// GetProject will get the project indicated by the key from the data store
func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	p := models.Project{
		Key: vars["key"],
	}

	err := Store.Projects().Get(&p)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}
}

// GetAllProjects will get all the projects on this instance that the user has
// permissions to
func GetAllProjects(w http.ResponseWriter, r *http.Request) {

}
