package main

import (
	"log"
	"net/http"

	"github.com/manigandand/angago/config"
)

var angagoConfig *AngaGoConf

func init() {
	angagoConfig = getConfigOrDefault()
}

func main() {
	port := ":" + config.Port
	http.HandleFunc("/", angagoServer)
	log.Println("angago listening on ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err.Error())
	}
}
