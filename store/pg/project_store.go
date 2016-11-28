package pg

import (
	"database/sql"
	"encoding/json"

	"github.com/praelatus/backend/models"
)

// ProjectStore contains methods for storing and retrieving Projects from a
// Postgres DB
type ProjectStore struct {
	db *sql.DB
}

func IntoProject(row rowScanner, p *models.Project) error {
	var lead models.User
	var ljson json.RawMessage

	err := row.Scan(&p.ID, &p.CreatedDate, &p.Name, &p.Key,
		&p.Homepage, &p.IconURL, &p.Repo, &ljson)
	if err != nil {
		return err
	}

	err = json.Unmarshal(ljson, &lead)
	p.Lead = lead

	return err
}

// Get gets a project by it's ID in a postgres DB.
func (ps *ProjectStore) Get(p *models.Project) error {
	var row *sql.Row

	switch p.Key {
	case "":
		row = ps.db.QueryRow(`SELECT id, created_date, name, 
								   key, homepage, icon_url, repo,
								   rows_to_json(lead.*)
							FROM projects 
							JOIN users AS lead ON lead.id = projects.lead_id
							WHERE id = $1;`, p.ID)
	default:
		row = ps.db.QueryRow(`SELECT id, created_date, name, 
								   key, homepage, icon_url, repo,
								   rows_to_json(lead.*)
							FROM projects 
							JOIN users AS lead ON lead.id = projects.lead_id
							WHERE key = $1;`, p.Key)

	}

	err := IntoProject(row, p)
	return handlePqErr(err)
}

// GetByKey gets a project by it's project key
func (ps *ProjectStore) GetByKey(team models.Team, p *models.Project) error {
	row := ps.db.QueryRow(`SELECT p.id, p.created_date, p.name, 
								  p.key, p.repo, p.homepage, 
								  p.icon_url, p.lead_id, p.team_id,
								  row_to_json(lead.*)
						    FROM projects AS p
							JOIN users AS lead ON lead.id = p.lead_id
							WHERE p.key = $1`, p.Key)

	err := IntoProject(row, p)
	return handlePqErr(err)
}

// GetAll returns all projects
func (ps *ProjectStore) GetAll() ([]models.Project, error) {
	var projects []models.Project

	rows, err := ps.db.Query(`SELECT p.id, p.created_date, p.name, 
								  p.key, p.repo, p.homepage, 
								  p.icon_url, p.lead_id, p.team_id,
								  row_to_json(lead.*)
							  FROM projects;`)
	if err != nil {
		return projects, handlePqErr(err)
	}

	for rows.Next() {
		var p models.Project

		err = IntoProject(rows, &p)
		if err != nil {
			return projects, handlePqErr(err)
		}

		projects = append(projects, p)
	}

	return projects, nil
}

// New creates a new Project in the database.
func (ps *ProjectStore) New(project *models.Project) error {
	err := ps.db.QueryRow(`INSERT INTO projects 
						   (name, key, repo, homepage, icon_url, lead_id) 
						   VALUES ($1, $2, $3, $4, $5, $6)
						   RETURNING id;`,
		project.Name, project.Key, project.Repo, project.Homepage,
		project.IconURL, project.Lead.ID).
		Scan(&project.ID)

	return handlePqErr(err)
}

// Save updates a Project in the database.
func (ps *ProjectStore) Save(project models.Project) error {
	_, err := ps.db.Exec(`UPDATE projects SET
						  (name, key, repo, homepage, icon_url, lead_id) 
						  = ($1, $2, $3, $4, $5, $6)
						  WHERE id = $7;`,
		project.Name, project.Key, project.Repo, project.Homepage,
		project.IconURL, project.Lead.ID, project.ID)

	return handlePqErr(err)
}

// Remove updates a Project in the database.
func (ps *ProjectStore) Remove(project models.Project) error {
	_, err := ps.db.Exec(`
	DELETE FROM field_tickettype_project WHERE project_id = $1;
	DELETE FROM permissions WHERE project_id = $1;
	DELETE FROM field_values
		WHERE ticket_id in(SELECT id FROM tickets WHERE project_id = $1);
	DELETE FROM ticket_labels 
		WHERE ticket_id in(SELECT id FROM tickets WHERE project_id = $1);
	DELETE FROM tickets WHERE project_id = $1;
	DELETE FROM projects WHERE id = $1;`, project.ID)

	return handlePqErr(err)
}
