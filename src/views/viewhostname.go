package views

import (
	"log"
	"net/http"
	"os"
)

// ViewHostname return Hostname
func ViewHostname(writer http.ResponseWriter, request *http.Request) {
	message, _ := (os.Hostname())
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
