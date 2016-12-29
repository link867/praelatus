package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/backend/models"
)

func TestGetType(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/types/1", nil)

	Router.ServeHTTP(w, r)

	var p models.TicketType

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p.ID != 1 {
		t.Errorf("Expected 1 Got %d\n", p.ID)
	}

	t.Log(w.Body)
}

func TestGetAllTypes(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/types", nil)
	testLogin(r)

	Router.ServeHTTP(w, r)

	var p []models.TicketType

	e := json.Unmarshal(w.Body.Bytes(), &p)
	if e != nil {
		t.Errorf("Failed with error %s\n", e.Error())
	}

	if p[0].Name != "mock type" {
		t.Errorf("Expected mock type Got %s\n", p[0].Name)
	}

	if len(p) != 2 {
		t.Errorf("Expected 2 Got %d\n", len(p))
	}

	t.Log(w.Body)
}

func TestCreateType(t *testing.T) {
	p := models.TicketType{Name: "Snug"}
	byt, _ := json.Marshal(p)
	rd := bytes.NewReader(byt)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/types", rd)
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
