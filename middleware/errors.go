package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	// ErrForbidden is used when the user does not have admin permissions
	ErrForbidden = errors.New("You do not have permission to access that resource")
	// ErrUnauthorized is used when the user is not logged in
	ErrUnauthorized = errors.New("You need to be logged in to access that resource")
)

func middlewareErr(w http.ResponseWriter, r *http.Request, e error, s time.Time) {
	var status int

	if ErrForbidden == e {
		status = 403
	}

	if ErrUnauthorized == e {
		status = 401
	}

	w.WriteHeader(status)
	w.Write([]byte(e.Error()))
	log.Printf("|%s| [%d] %s %s",
		r.Method, status, r.URL.Path, time.Since(s).String())
}
