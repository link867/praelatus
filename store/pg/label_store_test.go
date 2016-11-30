package pg_test

import (
	"testing"

	"github.com/praelatus/backend/models"
)

func TestLabelGet(t *testing.T) {
	l, e := s.Labels().Get(&models.Label{ID: 1})
	failIfErr("Label Get", t, e)

	if l == nil {
		t.Error("Expected a label and got nil instead.")
	}

	if l.Name == "" {
		t.Errorf("Expected label to have name got %s\n", l.Name)
	}
}

func TestLabelGetAll(t *testing.T) {
	l, e := s.Labels().GetAll()
	failIfErr("Label Get All", t, e)

	if l == nil {
		t.Error("Expected to get some labels and got nil instead.")
	}

	if len(l) == 0 {
		t.Errorf("Expected labels to be returned got %d labels instead\n", len(l))
	}
}

func TestLabelSave(t *testing.T) {
	l, e := s.Labels().Get(&models.Label{ID: 1})
	failIfErr("Label Save", t, e)

	l.Name = "SAVE_TEST_LABEL"

	e = s.Labels().Save(l)
	failIfErr("Label Save", t, e)

	l, e = s.Labels().Get(&models.Label{ID: 1})
	failIfErr("Label Save", t, e)

	if l.Name != "SAVE_TEST_LABEL" {
		t.Errorf("Expected: SAVE_TEST_LABEL Got: %s\n", l.Name)
	}
}

func TestLabelRemove(t *testing.T) {
	l, e := s.Labels().Get(&models.Label{ID: 2})
	failIfErr("Label Remove", t, e)

	e = s.Labels().Remove(l)
	failIfErr("Label Remove", t, e)
}
