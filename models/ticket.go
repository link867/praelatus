package models

import "time"

// TicketType represents the type of ticket.
type TicketType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Ticket represents a ticket
type Ticket struct {
	ID          int64        `json:"id"`
	CreatedDate time.Time    `json:"created_date"`
	UpdatedDate time.Time    `json:"updated_date"`
	Key         string       `json:"key"`
	Summary     string       `json:"summary"`
	Description string       `json:"description"`
	Fields      []FieldValue `json:"fields"`
	Labels      []Label      `json:"labels"`
	Type        TicketType   `json:"ticket_type"`
	Reporter    User         `json:"reporter"`
	Assignee    User         `json:"assignee"`
	Status      Status       `json:"status"`

	Comments []Comment `json:"comments,omitempty"`
}

func (t *Ticket) String() string {
	return jsonString(t)
}

// Status represents a ticket's current status.
type Status struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *Status) String() string {
	return jsonString(s)
}

// Label is a label used on tickets
type Label struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (l *Label) String() string {
	return jsonString(l)
}

//Link represents tickets linked to one another
type Link struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	ParentID        int64  `json:"-"`
	ToTicketID      int64  `json:"ticket_id"`
	ToTicketSummary string `json:"ticket_summary"`
	ToTicketKey     string `json:"ticket_key"`
	Href            string `json:"href"`
}

func (l *Link) String() string {
	return jsonString(l)
}
