package api

// import (
// 	"encoding/json"
// 	"net/http"

// 	mw "github.com/praelatus/backend/middleware"
// )

// func InitProjectRoutes() {
// 	BaseRoutes.Projects.Handle("/", mw.Auth(ListProjects)).Methods("GET")
// 	BaseRoutes.Projects.Handle("/{team_slug}/{pkey}", mw.Auth(GetProject)).Methods("GET")
// }

// // TODO
// func ListProjects(c *mw.Context) (int, []byte) {
// 	projects, err := Store.Projects().GetAll()
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	jsn, err := json.Marshal(projects)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, jsn
// }

// // TODO
// func GetProject(c *mw.Context) (int, []byte) {
// 	p, err := Store.Projects().GetByKey(c.Var("team_slug"), c.Var("key"))

// 	// TODO: better error handling
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	if p.Key == "" {
// 		return http.StatusNotFound, []byte("Project does not exist.")
// 	}

// 	jsn, err := json.Marshal(&p)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, jsn
// }
