package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// ViewHostname return Hostname
func ViewHostname(writer http.ResponseWriter, request *http.Request) {
	message, _ := (os.Hostname())
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

// ViewHostname return local time
func ViewTime(writer http.ResponseWriter, request *http.Request) {
	today := []byte(time.Now().String())
	_, err := writer.Write(today)
	if err != nil {
		log.Fatal(err)
	}
}

// ViewHostname return IP
func ViewIP(writer http.ResponseWriter, request *http.Request) {
	message := []byte("IP: ")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				message = []byte(ipnet.IP.String())
			}
		}
	}

	_, er := writer.Write([]byte(message))
	if er != nil {
		log.Fatal(err)
	}
}

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
