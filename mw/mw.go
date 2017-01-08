package mw

import "net/http"

// Default is the default middleware stack for Praelatus
var Default = []func(http.Handler) http.Handler{
	Logger,
	Auth,
}
