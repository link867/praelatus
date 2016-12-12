package mw

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHttp(w, r)

		log.Printf("|%s| [%d] %s %s",
			r.Method, statusCode, r.URL.Path, time.Since(start).String())
	})
}
