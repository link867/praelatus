package migrations

import (
	"database/sql"
	"log"
)

type schema struct {
	v    int
	q    string
	name string
}

var schemas = []schema{
	v1schema,
	v2schema,
	v3schema,
	v4schema,
	v5schema,
	v6schema,
	v7schema,
	v8schema,
	v9schema,
}

// SchemaVersion will find the schema version for the given database
func SchemaVersion(db *sql.DB) int {
	var v int

	rw, err := db.Query("SELECT schema_version FROM database_information WHERE id = 1")
	if err != nil {
		return 0
	}
	defer rw.Close()

	rw.Next()
	err = rw.Scan(&v)
	if err != nil {
		return 0
	}

	return v

}

// RunMigrations will run all database migrations depending on the version
// returned from the database_information table.
func RunMigrations(db *sql.DB) error {
	version := SchemaVersion(db)

	for _, schema := range schemas {
		version = SchemaVersion(db)

		if version < schema.v {
			log.Printf("Migrating database to version %d: %s\n", schema.v, schema.name)
			_, err := db.Exec(schema.q)
			if err != nil {
				return err
			}

			_, err = db.Exec(`UPDATE database_information 
							  SET (schema_version) = ($1)
							  WHERE id = 1;`, schema.v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
