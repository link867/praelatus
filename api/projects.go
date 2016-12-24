package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/mw"
)

func initProjectRoutes() {
	Router.Handle("/projects", mw.Default(getAllProjects)).Methods("GET")
	Router.Handle("/projects", mw.Default(createProject)).Methods("POST")
	Router.Handle("/projects/{pkey}", mw.Default(getProject)).Methods("GET")
}

func getProject(w http.ResponseWriter, r *http.Request) {
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

	sendJson(w, p)
}

// getAllProjects will get all the projects on this instance that the user has
// permissions to
// TODO handler permissions
func getAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := Store.Projects().GetAll()
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJson(w, projects)
}

func createProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	u := mw.GetUser(r.Context())
	if u == nil || !u.IsAdmin {
		w.WriteHeader(403)
		w.Write(apiError("you must be logged in as a system administrator to create a project"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError("invalid body"))
		log.Println(err)
		return
	}

	err = Store.Projects().New(&p)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJson(w, p)
}
