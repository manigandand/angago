package main

import (
	"fmt"
	"net/http"
)

var angagoConfig *AngaGoConf

func init() {
	angagoConfig = getConfigOrDefault()
}

func main() {
	port := ":80"
	http.HandleFunc("/", angagoServer)
	fmt.Println("angago listening on ", port)
	http.ListenAndServe(port, nil)
}
