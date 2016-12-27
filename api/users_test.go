package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/backend/models"
)

func TestGetUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/foouser", nil)

	Router.ServeHTTP(w, r)

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
	r := httptest.NewRequest("GET", "/users", nil)
	testAdminLogin(r)

	Router.ServeHTTP(w, r)

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

	t.Log(w.Body)
}

func TestCreateUser(t *testing.T) {
	u := models.User{Username: "grumpycat"}
	byt, _ := json.Marshal(u)
	rd := bytes.NewReader(byt)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users", rd)

	Router.ServeHTTP(w, r)

	var l TokenResponse

	e := json.Unmarshal(w.Body.Bytes(), &l)
	if e != nil {
		t.Errorf("Failed with error %s", e.Error())
	}

	if l.User.ID != 1 {
		t.Errorf("Expected 1 Got %d", u.ID)
	}

	if l.Token == "" {
		t.Errorf("Expected a token got %s\n", l.Token)
	}

	t.Log(w.Body)
}

func TestRefreshSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/sessions", nil)
	testLogin(r)

	Router.ServeHTTP(w, r)

	if w.Body.String() == "" {
		t.Errorf("Expected a token response got %s\n", w.Body.String())
	}

	if w.Code != 200 {
		t.Errorf("Expected 200 Got %d\n", w.Code)
	}

	t.Log(w.Body)
}
