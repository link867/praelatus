package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/praelatus/models"
)

func TestGetUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users/foouser", nil)

	router.ServeHTTP(w, r)

	var u models.User

	e := json.Unmarshal(w.Body.Bytes(), &u)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if u.Username != "foouser" {
		t.Errorf("Expected foouser Got %s\n", u.Username)
	}

	if u.Password != "" {
		t.Error("Expected no password to be returned but instead got a password.")
	}

	t.Log(w.Body)
}

func TestGetAllUsers(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users", nil)
	testAdminLogin(r)

	router.ServeHTTP(w, r)

	var u []models.User

	e := json.Unmarshal(w.Body.Bytes(), &u)
	if e != nil {
		t.Errorf("Failed with error %s", e.Error())
		t.Log(w.Body)
	}

	if len(u) != 2 {
		t.Errorf("Expected 2 users got %d", len(u))
	}

	if u[0].Username != "foouser" {
		t.Errorf("Expected foouser Got %s", u[0].Username)
	}

	if u[0].Password != "" {
		t.Errorf("Expected no passsword but got %s\n", u[0].Password)
	}

	for _, usr := range u {
		t.Log(usr.String())
	}
}

func TestCreateUser(t *testing.T) {
	u := models.User{Username: "grumpycat"}
	byt, _ := json.Marshal(u)
	rd := bytes.NewReader(byt)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/v1/users", rd)

	router.ServeHTTP(w, r)

	var l models.User

	e := json.Unmarshal(w.Body.Bytes(), &l)
	if e != nil {
		t.Errorf("Failed with error %s", e.Error())
	}

	if l.ID != 1 {
		t.Errorf("Expected 1 Got %d", u.ID)
	}

	if l.ProfilePic == "" {
		t.Error("Expected a profile pic but got nothing.")
	}

	t.Log(w.Body)
}

func TestRefreshSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users/sessions", nil)
	testLogin(r)

	router.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Expected 200 Got %d\n", w.Code)
	}

	t.Log(w.Body)
}

func TestSearchUsers(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users/search?query=foo", nil)
	testAdminLogin(r)

	router.ServeHTTP(w, r)

	var u []models.User

	e := json.Unmarshal(w.Body.Bytes(), &u)
	if e != nil {
		t.Errorf("Failed with error %s", e.Error())
		t.Log(w.Body)
	}

	if len(u) != 2 {
		t.Errorf("Expected 2 users got %d", len(u))
	}

	if u[0].Username != "foouser" {
		t.Errorf("Expected foouser Got %s", u[0].Username)
	}

	if u[0].Password != "" {
		t.Errorf("Expected no passsword but got %s\n", u[0].Password)
	}
}
