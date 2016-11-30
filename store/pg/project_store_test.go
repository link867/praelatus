package pg_test

import (
	"testing"

	"github.com/praelatus/backend/models"
)

func TestProjectGet(t *testing.T) {
	p := &models.Project{ID: 1}
	e := s.Projects().Get(p)
	failIfErr("Project Get", t, e)

	if p.Key == "" {
		t.Errorf("Expected: TEST Got: %s\n", p.Key)
	}

	p1 := &models.Project{Key: "TEST"}
	e := s.Projects().Get(p1)
	failIfErr("Project Get", t, e)

	if p1.ID == 0 {
		t.Errorf("Expected: 1 Got: %d\n", p1.ID)
	}
}

func TestProjectGetAll(t *testing.T) {
	p, e := s.Projects().GetAll()
	failIfErr("Project Get All", t, e)

	if p == nil || len(p) == 0 {
		t.Error("Expected to get some projects and got nil instead.")
	}
}

func TestProjectSave(t *testing.T) {
	p := &models.Project{ID: 1}
	e := s.Projects().Get(p)
	failIfErr("Project Save", t, e)

	p.IconURL = "TEST"

	e = s.Projects().Save(p)
	failIfErr("Project Save", t, e)

	p1 := &models.Project{ID: 1}
	e = s.Projects().Get(p1)
	failIfErr("Project Save", t, e)

	if p1.IconURL != "TEST" {
		t.Errorf("Expected project to have iconURL TEST got %s\n", p.IconURL)
	}
}

func TestProjectRemove(t *testing.T) {
	p := &models.Project{ID: 2}
	e := s.Projects().Remove(p)
	failIfErr("Project Remove", t, e)

	e = s.Projects().Get(p)
	failIfErr("Project Remove", t, e)

	if p.Key != "" {
		t.Errorf("Expected: \"\" Got :%s\n", p.Key)
	}
}
