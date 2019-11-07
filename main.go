package main

import (
	"fmt"
	"net/http"
)

func init() {

}

func main() {
	port := ":80"
	http.HandleFunc("/", angagoServer)
	fmt.Println("angago listening on ", port)
	http.ListenAndServe(port, nil)
}
