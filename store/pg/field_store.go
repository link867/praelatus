package pg

import (
	"database/sql"
	"errors"

	"github.com/praelatus/backend/models"
)

// FieldStore contains methods for storing and retrieving Fields and
// FieldValues in a Postgres Database
type FieldStore struct {
	db *sql.DB
}

// Get retrieves a models.Field by ID
func (fs *FieldStore) Get(f *models.Field) error {
	var row *sql.Row

	row = fs.db.QueryRow(`SELECT id, name, data_type FROM fields 
						  WHERE id = $1 OR name = $2`, f.ID, f.Name)
	err := row.Scan(&f.ID, &f.Name, &f.DataType)

	return handlePqErr(err)
}

// GetAll will return all fields from the DB
func (f *FieldStore) GetAll() ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Query("SELECT id, name, data_type FROM fields;")
	if err != nil {
		return fields, handlePqErr(err)
	}

	for rows.Next() {
		var f models.Field

		err = rows.Scan(&f.ID, &f.Name, &f.DataType)
		if err != nil {
			return fields, handlePqErr(err)
		}

		fields = append(fields, f)
	}

	return fields, nil
}

// GetByProject retrieves all Fields associated with a project
func (f *FieldStore) GetByProject(p models.Project) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Query(`
		SELECT fields.id, fields.name, fields.data_type FROM fields
		JOIN field_tickettype_project as ftp ON fields.id = ftp.field_id
		WHERE ftp.key = $1;`, p.Key)
	if err != nil {
		return fields, handlePqErr(err)
	}

	for rows.Next() {
		var f models.Field

		err = rows.Scan(&f.ID, &f.Name, &f.DataType)
		if err != nil {
			return fields, handlePqErr(err)
		}

		fields = append(fields, f)
	}

	return fields, nil
}

// AddToProject adds a field to a project's tickets
func (f *FieldStore) AddToProject(project models.Project, field *models.Field,
	ticketTypes ...models.TicketType) error {

	if ticketTypes == nil {
		_, err := f.db.Exec(`INSERT INTO field_tickettype_project 
							 (field_id, project_id) VALUES ($1, $2)`,
			field.ID, project.ID)
		return handlePqErr(err)
	}

	for _, typ := range ticketTypes {

		_, err := f.db.Exec(`INSERT INTO field_tickettype_project 
							 (field_id, project_id, ticket_type_id) 
							 VALUES ($1, $2, $3)`,
			field.ID, project.ID, typ.ID)
		if err != nil {
			return handlePqErr(err)
		}
	}

	return nil

}

// Save updates an existing field in the database.
func (f *FieldStore) Save(field models.Field) error {
	_, err := f.db.Exec(`UPDATE fields SET 
					     (name, data_type) = ($1, $2) WHERE id = $3;`,
		field.Name, field.DataType, field.ID)

	return handlePqErr(err)
}

// New creates a new Field in the database.
func (f *FieldStore) New(field *models.Field) error {
	err := f.db.QueryRow(`INSERT INTO fields 
						  (name, data_type) VALUES ($1, $2)
						  RETURNING id;`,
		field.Name, field.DataType).
		Scan(&field.ID)

	return handlePqErr(err)
}

// Remove updates an existing field in the database.
func (f *FieldStore) Remove(field models.Field) error {
	var c int

	err := f.db.QueryRow(`SELECT COUNT(id) FROM field_values 
						  WHERE field_id = $1`, field.ID).Scan(&c)
	if err != nil {
		return handlePqErr(err)
	}

	if c > 0 {
		return errors.New("that field is currently in use, refusing to delete")
	}

	_, err = f.db.Exec("DELETE FROM fields WHERE id = $1", field.ID)
	return handlePqErr(err)
}
