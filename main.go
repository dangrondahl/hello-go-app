package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello world received a request.")
	fmt.Fprintf(w, "Hello World: Example!\n")
}

func main() {
	log.Printf("Hello world sample started. Version 1")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
