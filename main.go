package main

import (
	"log"
	"net/http"
)

var angagoConfig *AngaGoConf

func init() {
	angagoConfig = getConfigOrDefault()
}

func main() {
	port := ":80"
	http.HandleFunc("/", angagoServer)
	log.Println("angago listening on ", port)
	http.ListenAndServe(port, nil)
}
