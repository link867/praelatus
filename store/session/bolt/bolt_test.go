package bolt

import (
	"testing"

	"github.com/praelatus/praelatus/models"
	"github.com/praelatus/praelatus/store"
)

var s store.SessionStore

func init() {
	s = New("SessionTest.db")
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
