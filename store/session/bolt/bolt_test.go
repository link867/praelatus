package bolt

import (
	"testing"

	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

var s store.SessionStore

func init() {
	var err error
	s, err = New("SessionTest.db")
	if err != nil {
		panic(err)
	}
}

func TestGetAndSet(t *testing.T) {
	user, _ := models.NewUser("testuser", "test", "Test Testerson", "test@example.com", false)

	err := s.Set("test", *user)
	if err != nil {
		t.Error(err)
	}

	u, err := s.Get("test")
	if err != nil {
		t.Error(err)
	}

	if u != *user {
		t.Errorf("Expected %v, got %v", user, u)
	}
}
