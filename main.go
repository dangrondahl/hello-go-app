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
		"Starting the service...\ncommit: %s, build time: %s",
		version.Commit, version.BuildTime)
	router := handlers.Router(version.BuildTime, version.Commit)
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
