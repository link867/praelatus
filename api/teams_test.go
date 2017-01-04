package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/backend/models"
)

func TestGetTeam(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/teams/1", nil)

	router.ServeHTTP(w, r)

	var p models.Team

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p.ID != 1 {
		t.Errorf("Expected 1 Got %d\n", p.ID)
	}

	t.Log(w.Body)
}

func TestGetAllTeams(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/teams", nil)
	testLogin(r)

	router.ServeHTTP(w, r)

	var p []models.Team

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p[0].Name != "A" {
		t.Errorf("Expected A Got %s\n", p[0].Name)
	}

	if len(p) != 2 {
		t.Errorf("Expected 2 Got %d\n", len(p))
	}

	t.Log(w.Body)
}

func TestCreateTeam(t *testing.T) {
	p := models.Team{Name: "Snug"}
	byt, _ := json.Marshal(p)
	rd := bytes.NewReader(byt)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/teams", rd)
	testAdminLogin(r)

	router.ServeHTTP(w, r)

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s", e.Error())
	}

	if p.ID != 1 {
		t.Errorf("Expected 1 Got %d", p.ID)
	}

	t.Log(w.Body)
}
