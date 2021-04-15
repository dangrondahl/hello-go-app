package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dangrondahl/hello-go-app/version"
)

func root(w http.ResponseWriter, r *http.Request) {
	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
	}{
		version.BuildTime, version.Commit,
	}

	body, err := json.Marshal(info)

	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	log.Printf("Hello world received a request.")
}
