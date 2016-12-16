package api

// import (
// 	"encoding/json"
// 	"net/http"

// 	mw "github.com/praelatus/backend/middleware"
// 	"github.com/praelatus/backend/models"
// )

// func InitTicketRoutes() {
// 	BaseRoutes.Tickets.Handle("/{team_slug}/{pkey}/{key}", mw.Default(GetTicket)).Methods("GET")
// 	BaseRoutes.Tickets.Handle("/{team_slug}/{pkey}/{key}", mw.Default(CreateTicket)).Methods("POST")
// 	BaseRoutes.Tickets.Handle("/{team_slug}/{pkey}/{key}", mw.Default(UpdateTicket)).Methods("PUT")
// }

// // TODO: Fix error handling

// // ListTickets will list all tickets in the database
// func ListTickets(c *mw.Context) (int, []byte) {
// 	tickets, err := Store.Tickets().GetAll()
// 	if err != nil {
// 		return 500, []byte(err.Error())
// 	}

// 	tjson, err := json.Marshal(tickets)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, tjson
// }

// // CreateTicket creates a ticket
// // TODO
// func CreateTicket(c *mw.Context) (int, []byte) {
// 	var t models.TicketJSON
// 	err := c.JSON(&t)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	err = Store.Tickets().Save(models.TicketFromJSON(t))
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, []byte("Ticket successfully created.")
// }

// // UpdateTicket will update the ticket at :key
// // TODO
// func UpdateTicket(c *mw.Context) (int, []byte) {
// 	return http.StatusNotImplemented, []byte("Not implemented")
// }

// // GetTicket will get a specific ticket
// func GetTicket(c *mw.Context) (int, []byte) {
// 	key := c.Var("key")
// 	pkey := c.Var("pkey")
// 	teamSlug := c.Var("team_slug")

// 	t, err := Store.Tickets().GetByKey(teamSlug, pkey, key)
// 	if err != nil {
// 		return http.StatusNotFound, []byte(err.Error())
// 	}

// 	jsn, err := json.Marshal(&t)
// 	if err != nil {
// 		return http.StatusInternalServerError, []byte(err.Error())
// 	}

// 	return http.StatusOK, jsn
// }

// // TODO: Implement workflows
// // AdvanceTicket advances a ticket through it's workflow.
// func AdvanceTicket(c *mw.Context) (int, []byte) {
// 	return http.StatusNotImplemented, []byte("Not implemented")
// }

// // RetractTicket will move a ticket backwards in it's workflow
// func RetractTicket(c *mw.Context) (int, []byte) {
// 	return http.StatusNotImplemented, []byte("Not implemented")
// }
