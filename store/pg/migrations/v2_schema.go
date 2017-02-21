package migrations

const addTicketLinks = `
CREATE TABLE IF NOT EXISTS ticket_links (
	id SERIAL PRIMARY KEY,
	name varchar(40) NOT NULL,
	parent_ticket integer REFERENCES tickets (id),
	linked_ticket integer REFERENCES tickets (id)
);`

var v10schema = schema{10, addTicketLinks, "add ticket links"}
