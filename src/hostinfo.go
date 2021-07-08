package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"views"
)

func main() {

	var (
		port string = "8080"
		ip   string = "0.0.0.0"
	)

	io.WriteString(os.Stdout, "Run service on "+ip+":"+port+"\n")

	http.HandleFunc("/", views.ViewHostname)
	http.HandleFunc("/ip", views.ViewIP)
	http.HandleFunc("/time", views.ViewTime)

	err := http.ListenAndServe(ip+":"+port, nil)
	log.Fatal(err)

}
