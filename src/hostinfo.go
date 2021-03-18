package main

import (
	"fmt"
	"log"
	"net/http"
	"views"
)

func main() {

	var (
		port string = "8080"
		ip   string = "127.0.0.1"
	)

	fmt.Println("Run service on " + ip + ":" + port)

	http.HandleFunc("/", views.ViewHostname)
	http.HandleFunc("/ip", views.ViewIP)
	http.HandleFunc("/time", views.ViewTime)

	err := http.ListenAndServe(ip+":"+port, nil)
	log.Fatal(err)

}
