package handlers

import (
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit string) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(3 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", root(buildTime, commit)).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")
	r.HandleFunc("/readyz", readyz(isReady)).Methods("GET")
	return r
}
