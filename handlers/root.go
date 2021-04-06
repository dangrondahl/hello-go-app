package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello world received a request.")
	fmt.Fprint(w, "Hello World: Example!")
}
