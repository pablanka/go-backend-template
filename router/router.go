package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Route defines a route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	IsStatic    bool
}

//Routes is an array of routes
type Routes []*Route

type mytype string

// wrapLog wraps a log on route's handler func
func logWrapper(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// NewRouter creates a new Router with all cofigured routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes { // creates a router for each route.
		if route.IsStatic {
			router.
				Methods(route.Method).
				PathPrefix(route.Pattern).
				Name(route.Name).
				Handler(logWrapper(route.HandlerFunc, route.Name))
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(logWrapper(route.HandlerFunc, route.Name))
		}
	}
	return router
}
