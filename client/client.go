package main

import (
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Info("%s\n", string(contents))
	}
}
