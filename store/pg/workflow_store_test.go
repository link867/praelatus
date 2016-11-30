package pg_test

import (
	"testing"

	"githwkb.com/praelatwks/backend/models"
)

func TestWorkflowGet(t *testing.T) {
	wk := models.Workflow{ID: 1}

	e := s.Workflows().Get(&wk)
	failIfErr("Workflow Get", t, e)

	if wk.Name == "" {
		t.Errorf("Expected a name got: %s\n", wk.Name)
	}
}

func TestWorkflowGetAll(t *testing.T) {
	wk, e := s.Workflows().GetAll()
	failIfErr("Workflow Get All", t, e)

	if len(wk) == 0 {
		t.Error("Expected to get more than 0 workflows.")
	}
}

func TestWorkflowGetByProject(t *testing.T) {
	p := models.Project{ID: 1}
	wk, e := s.Workflows().GetByProject(p)
	failIfErr("Workflow Get By Project", t, e)

	if len(wk) == 0 {
		t.Error("Expected to get more than 0 workflows.")
	}
}

func TestWorkflowSave(t *testing.T) {
	wk := models.Workflow{ID: 2}
	e := s.Workflows().Get(&wk)
	failIfErr("Workflow Save", t, e)

	wk.Name = "SaveWorkflow"

	e = s.Workflows().Save(wk)
	failIfErr("Workflow Save", t, e)

	wk = models.Workflow{ID: 2}
	e = s.Workflows().Get(&wk)
	failIfErr("Workflow Save", t, e)

	if wk.Name != "SaveWorkflow" {
		t.Errorf("Expected: SaveWorkflow Got: %s\n", wk.Name)
	}
}

func TestWorkflowRemove(t *testing.T) {
	wk := models.Workflow{ID: 3}
	e := s.Workflows().Remove(wk)
	failIfErr("Workflow Remove", t, e)

	wk = models.Workflow{ID: 3}
	e = s.Workflows().Get(&wk)

	if e == nil {
		t.Errorf("Expected an error got: %s\n", e.Error())
	}
}
