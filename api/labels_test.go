package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/backend/models"
)

func TestGetLabel(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/labels/1", nil)

	router.ServeHTTP(w, r)

	var p models.Label

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p.ID != 1 {
		t.Errorf("Expected 1 Got %d\n", p.ID)
	}

	t.Log(w.Body)
}

func TestGetAllLabels(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/labels", nil)
	testLogin(r)

	router.ServeHTTP(w, r)

	var p []models.Label

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p[0].Name != "mock" {
		t.Errorf("Expected mock Got %s\n", p[0].Name)
	}

	if len(p) != 2 {
		t.Errorf("Expected 2 Got %d\n", len(p))
	}

	t.Log(w.Body)
}

func TestCreateLabel(t *testing.T) {
	p := models.Label{Name: "Snug"}
	byt, _ := json.Marshal(p)
	rd := bytes.NewReader(byt)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/labels", rd)
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
