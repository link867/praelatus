package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

// LabelStore contains methods for storing and retrieving Labels from a
// Postgres DB
type LabelStore struct {
	db *sql.DB
}

// Get gets a label from the database by it's ID
func (ls *LabelStore) Get(ID int64) (*models.Label, error) {
	var l models.Label
	err := ls.db.QueryRowx("SELECT * FROM labels WHERE id = $1", ID).
		StructScan(&l)
	return &l, handlePqErr(err)
}

// GetAll gets all the labels from the database
func (ls *LabelStore) GetAll() ([]models.Label, error) {
	var labels []models.Label
	rows, err := ls.db.Queryx("SELECT * FROM labels;")

	for rows.Next() {
		var l models.Label

		err := rows.StructScan(&l)
		if err != nil {
			return labels, handlePqErr(err)
		}

		labels = append(labels, l)
	}

	return labels, handlePqErr(err)
}

// New creates a new label in the database
func (ls *LabelStore) New(label *models.Label) error {
	err := ls.db.QueryRow(`INSERT INTO labels (name) VALUES ($1)
						   RETURNING id;`,
		label.Name).
		Scan(&label.ID)

	return handlePqErr(err)
}

// Save updates a label in the database
func (ls *LabelStore) Save(label *models.Label) error {
	_, err := ls.db.Exec(`UPDATE labels SET (name) = ($1) 
						  WHERE id = $2;`, label.Name, label.ID)

	return handlePqErr(err)
}
