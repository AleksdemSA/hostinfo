package views

import (
	"log"
	"net"
	"net/http"
)

// ViewIP return IP of host
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
