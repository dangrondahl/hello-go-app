package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello world received a request.")
	fmt.Fprintf(w, "Hello World: Argo Example!\n")
}

func main() {
	log.Printf("Hello world sample started.")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
