package middleware

import (
	"net/http"
	"time"

	"log"
)

// ContextHandler is a custom function type which handles a http request, it takes
// a pointer to our context as the first argument and returns an integer and
// []byte which will be the response status code and body respectively.
type ContextHandler func(*Context) (int, []byte)

// Middleware is a function which takes a MiddlewareContext and returns a
// middleware.Context
type Middleware func(*Context) *Context

// Stack is a http.Handler which contains the ContextHandler and Middleware
// stack to run for a given request.
type Stack struct {
	Fn ContextHandler
	Mw []Middleware
}

// ServeHTTP implements the http.Handler interface for Stack.
func (s *Stack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	c := &Context{
		R:   r,
		Val: make(map[string]interface{}),
	}

	for _, mw := range s.Mw {
		c = mw(c)
		if c.Err != nil {
			middlewareErr(w, r, c.Err, start)
			return
		}
	}

	statusCode, response := s.Fn(c)

	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}

	if statusCode >= 300 {
		log.Printf("|%s| [%d] %s %s",
			r.Method, statusCode, r.URL.Path, time.Since(start).String())
	} else {
		log.Printf("|%s| [%d] %s %s",
			r.Method, statusCode, r.URL.Path, time.Since(start).String())
	}

}

var defaultMw = []Middleware{}

// Default returns our default middleware stack which requires no auth.
func Default(fn ContextHandler) http.Handler {
	return &Stack{
		Fn: fn,
		Mw: defaultMw,
	}
}

// Auth returns a stack with the default middleware plus the authentication
// middleware.
func Auth(fn ContextHandler) http.Handler {
	return &Stack{
		Fn: fn,
		Mw: append(defaultMw, AuthMw),
	}
}

// Admin returns a stack with the default middleware plus the authentication
// middleware.
func Admin(fn ContextHandler) http.Handler {
	return &Stack{
		Fn: fn,
		Mw: append(defaultMw, AdminMw),
	}
}
