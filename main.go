package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	path := "stub/default.conf.yaml"
	body, err := getBytesForFileOrURL(path)
	if err != nil {
		log.Fatal(err)
	}

	conf, err := decodeConfig(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("COnf\n%+v\n\n", conf)
}

func main() {
	port := ":80"
	http.HandleFunc("/", angagoServer)
	fmt.Println("angago listening on ", port)
	http.ListenAndServe(port, nil)
}
