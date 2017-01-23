package mw

import (
	"log"
	"net/http"

	"github.com/praelatus/backend/models"
)

// GetUser will check the given http.Request for a session token and if found
// it will return the corresponding user.
func GetUser(r *http.Request) *models.User {
	cookie, err := r.Cookie("PRASESSION")
	if err != nil {
		log.Println("Error getting cookie:", err)
		return nil
	}

	token := cookie.String()

	return nil
}
