package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/backend/models"
)

func TestGetProject(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/projects/TEST", nil)

	Router.ServeHTTP(w, r)

	var p models.Project

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p.Key != "TEST" {
		t.Errorf("Expected TEST-1 Got %s\n", p.Key)
	}

	t.Log(w.Body)
}

func TestGetAllProjects(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/projects", nil)

	Router.ServeHTTP(w, r)

	var p []models.Project

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p[0].Key != "TEST" {
		t.Errorf("Expected TEST-1 Got %s\n", p[0].Key)
	}

	if len(p) != 2 {
		t.Errorf("Expected 2 Got %d\n", len(p))
	}

	t.Log(w.Body)
}

func TestCreateProject(t *testing.T) {
	p := models.Project{Name: "Grumpy Cat", Key: "NOPE"}
	byt, _ := json.Marshal(p)
	rd := bytes.NewReader(byt)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/projects", rd)
	testAdminLogin(r)

	Router.ServeHTTP(w, r)

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s", e.Error())
	}

	if p.ID != 1 {
		t.Errorf("Expected 1 Got %d", p.ID)
	}

	t.Log(w.Body)
}
