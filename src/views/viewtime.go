package views

import (
	"log"
	"net/http"
	"time"
)

// ViewHostname return local time
func ViewTime(writer http.ResponseWriter, request *http.Request) {
	today := []byte(time.Now().String())
	_, err := writer.Write(today)
	if err != nil {
		log.Fatal(err)
	}
}
