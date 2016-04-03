package main

import (
	"io"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type serverConfig struct {
	port int
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Received request")
	io.WriteString(w, "Got ya!")
}

func main() {
	log.Info("Starting HTTP Server on port")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
