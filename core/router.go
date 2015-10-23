package core

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Router is gorilla's mux with an array of middlewares to go through
type Router struct {
	*mux.Router
	middlewares []Middleware
}

//Middleware holds a path and handler for a middleware
type Middleware struct {
	path    string
	handler http.Handler
}

//NewRouter reates a new router and returns it's pointer
func NewRouter() *Router {
	return &Router{mux.NewRouter(), []Middleware{}}
}

// AddMiddleware lets you inject middlewares into the routes
// It takes a regular string and a http.Handler. The string is the path which the middleware applies to and the Handler is
// the function to run for the middleware.
func (rt *Router) AddMiddleware(p string, h http.Handler) {
	rt.middlewares = append(rt.middlewares, Middleware{p, h})
}

// Override the default ServeHTTP to run middlewares first then the real Handler
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Loop through all the middlewares and run them if the path matches
	for key := range rt.middlewares {
		middleware := rt.middlewares[key]

		// Check if the path matches the path of the middleware and skip the for loop if it doesnt match
		if !strings.Contains(r.URL.Path, middleware.path) {
			continue
		}

		// Run the middleware function
		middleware.handler.ServeHTTP(w, r)
	}

	//Run the parent function
	rt.Router.ServeHTTP(w, r)
}
