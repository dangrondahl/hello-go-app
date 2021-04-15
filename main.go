package main

import (
	"log"
	"net/http"

	"github.com/dangrondahl/hello-go-app/handlers"
	"github.com/dangrondahl/hello-go-app/version"
)

// How to try it: go run main.go
func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, version: %s",
		version.Commit, version.BuildTime, version.Version)
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8085", router))
}
