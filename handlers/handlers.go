package handlers

import (
	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", root(buildTime, commit)).Methods("GET")
	return r
}
