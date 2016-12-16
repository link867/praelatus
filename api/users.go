package api

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	mw "github.com/praelatus/backend/middleware"
// 	"github.com/praelatus/backend/models"
// )

// func InitUserRoutes() {
// 	BaseRoutes.Users.Handle("/", mw.Admin(ListUsers)).Methods("GET")
// 	BaseRoutes.Users.Handle("/", mw.Default(CreateUser)).Methods("POST")
// 	BaseRoutes.Users.Handle("/{name}", mw.Auth(GetUser)).Methods("GET")
// 	BaseRoutes.Users.Handle("/{name}", mw.Auth(UpdateUser)).Methods("POST", "PUT")
// }

// func ListUsers(c *mw.Context) (int, []byte) {
// 	users, err := Store.Users().GetAll()
// 	// TODO: better error handling
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	jsn, err := json.Marshal(&users)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, jsn
// }

// func CreateUser(c *mw.Context) (int, []byte) {
// 	// TODO: Handle errors properly (i.e. username taken etc.)
// 	var u models.User

// 	err := c.JSON(&u)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	err = Store.Users().Save(&u)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	jsn, err := json.Marshal(&u)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, jsn
// }

// func LoginUser(c *mw.Context) (int, []byte) {
// 	return http.StatusNotImplemented, []byte("Not implemented")
// }

// func GetUser(c *mw.Context) (int, []byte) {
// 	var u *models.User
// 	var err error

// 	username := c.Var("username")

// 	// If we get an integer search based on ID instead of username
// 	if id, err := strconv.Atoi(username); err == nil {
// 		u, err = Store.Users().Get(int64(id))
// 	} else {
// 		u, err = Store.Users().GetByUsername(username)
// 	}

// 	// TODO: Properly return not found when appropriate
// 	if err != nil {
// 		// TODO: Don't return raw terrible stuff.
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	ujson, err := json.Marshal(&u)

// 	return http.StatusOK, ujson
// }

// // TODO: implement this
// func UpdateUser(c *mw.Context) (int, []byte) {
// 	return http.StatusNotImplemented, []byte("Not implemented")
// }
