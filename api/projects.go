package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/praelatus/backend/models"
	"github.com/pressly/chi"
)

func projectRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", GetAllProjects)
	router.Post("/", CreateProject)

	router.Get("/:pkey", GetProject)
	router.Delete("/:pkey", RemoveProject)
	router.Put("/:pkey", UpdateProject)

	return router
}

// GetProject will get a project by it's project pkey
func GetProject(w http.ResponseWriter, r *http.Request) {
	pkey := chi.URLParam(r, "pkey")

	p := models.Project{
		Key: pkey,
	}

	err := Store.Projects().Get(&p)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, p)
}

// GetAllProjects will get all the projects on this instance that the user has
// permissions to
// TODO handle permissions
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	u := GetUserSession(r)
	if u == nil {
		w.WriteHeader(403)
		w.Write(apiError("you must be logged in to view all projects"))
		return
	}

	projects, err := Store.Projects().GetAll()
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, projects)
}

// CreateProject will create a project based on the JSON representation sent to
// the API
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	u := GetUserSession(r)
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

	sendJSON(w, p)
}

// RemoveProject will remove the project indicated by the pkey passed in as a
// url parameter
func RemoveProject(w http.ResponseWriter, r *http.Request) {
	pkey := chi.URLParam(r, "pkey")

	u := GetUserSession(r)
	if u == nil || !u.IsAdmin {
		w.WriteHeader(403)
		w.Write(apiError("you must be logged in as a system administrator to create a project"))
		return
	}

	err := Store.Projects().Remove(models.Project{Key: pkey})
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	w.Write([]byte{})

}

// UpdateProject will update a project based on the JSON representation sent to
// the API
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	u := GetUserSession(r)
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

	sendJSON(w, p)
}
