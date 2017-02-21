package pg

import (
	"database/sql"

	"github.com/praelatus/praelatus/models"
)

// LinkStore contains methods for storing and retrieving Links
type LinkStore struct {
	db *sql.DB
}

func intoLink(row rowScanner, l *models.Link) error {
	return row.Scan(&l.ID, &l.Name, &l.ParentID,
		&l.ToTicketID, &l.ToTicketSummary, &l.ToTicketKey)
}

// Get a Link from the database
func (ln *LinkStore) Get(l *models.Link) error {
	var row *sql.Row

	row = ln.db.QueryRow(`
		SELECT id, name, parent.id, 
			   linked.id, linked.summary, linked.key
		FROM ticket_links
		JOIN tickets as parent on parent.id = ticket_links.parent_ticket
		JOIN tickets as linked on linked.id = ticket_links.linked_ticket
		WHERE id = $1`, l.ID)

	return handlePqErr(intoLink(row, l))
}

// GetForTicket gets the links on a tickets
func (ln *LinkStore) GetForTicket(t models.Ticket) ([]models.Link, error) {

	rows, err := ln.db.Query(`
		SELECT id, name, parent.id, 
			   linked.id, linked.summary, linked.key
		FROM ticket_links
		JOIN tickets as parent on parent.id = ticket_links.parent_ticket
		JOIN tickets as linked on linked.id = ticket_links.linked_ticket
		WHERE parent.id = $1`, t.ID)
	if err != nil {
		return nil, handlePqErr(err)
	}

	var links = []models.Link{}

	for rows.Next() {
		var l models.Link

		err := intoLink(rows, &l)
		if err != nil {
			return links, handlePqErr(err)
		}

		links = append(links, l)
	}

	return links, nil
}

// New
func (ln *LinkStore) New(l *models.Link) error {

	err := ln.db.QueryRow(`
		INSERT INTO ticket_links 
		(name, parent_ticket, linked_ticket)
		VALUES ($1, $2, $3)
		RETURNING id;`,
		l.Name, l.ParentID, l.ToTicketID).
		Scan(&l.ID)

	return handlePqErr(err)

}

func (ln *LinkStore) Save(l models.Link) error {

	_, err := ln.db.Exec(`
		UPDATE ticket_links 
		(name, parent_ticket, linked_ticket)
		= ($1, $2, $3)
		WHERE id = $4;`,
		l.Name, l.ParentID, l.ToTicketID, l.ID)

	return handlePqErr(err)

}

func (ln *LinkStore) Remove(l models.Link) error {

	_, err := ln.db.Exec(`
		DELETE FROM ticket_links
		WHERE id = $1;`, l.ID)

	return handlePqErr(err)

}
