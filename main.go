package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	port := ":80"
	http.HandleFunc("/", angagoServer)
	fmt.Println("angago listening on ", port)
	http.ListenAndServe(port, nil)
}

// DomainRegistryMap ...
var DomainRegistryMap = map[string]string{
	"gopherhut.com":        "http://localhost:8080",
	"app.gopherhut.com":    "http://localhost:8081",
	"status.gopherhut.com": "http://localhost:8082",
	"admin.gopherhut.com":  "http://localhost:8083",
}

func angagoServer(w http.ResponseWriter, r *http.Request) {
	ReverseProxy(w, r, DomainRegistryMap[r.Host])
}

// ReverseProxy Serve a reverse proxy for a given url
func ReverseProxy(w http.ResponseWriter, r *http.Request, target string) {
	remote, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	// proxy.Transport = &transport{http.DefaultTransport}
	// proxy.ModifyResponse = modifyProxyResponse

	r.URL.Scheme = remote.Scheme
	r.URL.Host = remote.Host
	// r.URL.Path = path
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)

	return
}
