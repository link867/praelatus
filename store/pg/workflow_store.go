package pg

import (
	"database/sql"
	"encoding/json"

	"github.com/praelatus/backend/models"
)

// WorkflowStore contains methods for saving/retrieving workflows from a
// postgres DB
type WorkflowStore struct {
	db *sql.DB
}

// Get gets a workflow from the database
func (ws *WorkflowStore) Get(p models.Project, w *models.Workflow) error {
	row := ws.db.QueryRow(`SELECT id, name 
						   FROM workflows 
						   JOIN projects AS p ON workflows.project_id = p.id
						   WHERE id = $1 OR 
						   (name = $2 AND p.key)`, w.ID, w.Name, p.Key)

	err := row.Scan(&w.ID, &w.Name)
	if err != nil {
		return handlePqErr(err)
	}

	err = ws.getTransitions(w)
	return handlePqErr(err)
}

func (ws *WorkflowStore) getHooks(t *models.Transition) error {
	rows, err := ws.db.Query(`SELECT h.id, endpoint, method, body
						   FROM hooks AS h
						   JOIN transitions AS t ON t.id = h.transition_id`)
	if err != nil {
		return err
	}

	for rows.Next() {
		var h models.Hook

		err = rows.Scan(&h.ID, &h.Endpoint, &h.Method, &h.Body)
		if err != nil {
			return err
		}

		t.Hooks = append(t.Hooks, h)
	}

	return nil
}

func (ws *WorkflowStore) getTransitions(w *models.Workflow) error {
	rows, err := ws.db.Query(`SELECT t.id, t.name, 
									 from_s.name,
									 row_to_json(to_s.*)
							  FROM transitions AS t
							  JOIN statuses AS from_s ON from_s.id = t.from_status
							  JOIN statuses AS to_s ON to_s.id = t.status_id`)

	if err != nil {
		return handlePqErr(err)
	}

	for rows.Next() {
		var t models.Transition
		var fromStatus string
		var status json.RawMessage

		err = rows.Scan(&t.ID, &t.Name, &fromStatus, &status)
		if err != nil {
			return handlePqErr(err)
		}

		err = json.Unmarshal(status, &t.ToStatus)
		if err != nil {
			return handlePqErr(err)
		}

		err = ws.getHooks(&t)
		if err != nil {
			return handlePqErr(err)
		}

		w.Transitions[fromStatus] = append(w.Transitions[fromStatus], t)
	}

	return nil
}

func workflowsFromRows(rows *sql.Rows, ws *WorkflowStore) ([]models.Workflow, error) {
	var workflows []models.Workflow

	for rows.Next() {
		var w models.Workflow

		err := rows.Scan(&w.ID, &w.Name)
		if err != nil {
			return workflows, handlePqErr(err)
		}

		err = ws.getTransitions(&w)
		if err != nil {
			return workflows, handlePqErr(err)
		}

		workflows = append(workflows, w)
	}

	return workflows, nil
}

// GetAll gets all the workflows from the database
func (ws *WorkflowStore) GetAll() ([]models.Workflow, error) {
	rows, err := ws.db.Query("SELECT id, name FROM workflows;")
	if err != nil {
		return nil, handlePqErr(err)
	}

	return workflowsFromRows(rows, ws)
}

// GetByProject gets all the workflows for the given project
func (ws *WorkflowStore) GetByProject(p models.Project) ([]models.Workflow, error) {
	rows, err := ws.db.Query(`SELECT w.id, w.name 
							  FROM workflows AS w
							  JOIN projects AS p ON p.id = w.project_id
							  WHERE p.id = $1
							  OR p.key = $2`, p.ID, p.Key)
	if err != nil {
		return []models.Workflow{}, handlePqErr(err)
	}

	return workflowsFromRows(rows, ws)
}

// TODO handle transitions and hooks here

// New creates a new workflow in the database
func (ws *WorkflowStore) New(p models.Project, workflow *models.Workflow) error {
	err := ws.db.QueryRow(`INSERT INTO workflows 
						   (name, project_id) VALUES ($1, $2)
						   RETURNING id;`,
		workflow.Name, p.ID).
		Scan(&workflow.ID)

	return handlePqErr(err)
}

// Save updates a workflow in the database
func (ws *WorkflowStore) Save(w models.Workflow) error {
	_, err := ws.db.Exec(`UPDATE workflows SET (name) = ($1)`, w.Name)
	return handlePqErr(err)
}

// Remove removes a workflow from the database
func (ws *WorkflowStore) Remove(w models.Workflow) error {
	_, err := ws.db.Exec(`
	DELETE FROM hooks 
		WHERE transition_id in(SELECT id FROM transitions WHERE workflow_id = $1);
	DELETE FROM transitions WHERE workflow_id = $1;
	DELETE FROM workflows WHERE id = $1;`, w.ID)
	return handlePqErr(err)
}
