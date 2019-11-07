package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func angagoServer(w http.ResponseWriter, r *http.Request) {
	targetHost, ok := DomainRegistryMap[r.Host]
	if !ok {

	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	fmt.Println("tls", r.TLS)
	fmt.Printf("%s%s\n", scheme, r.URL.Host)
	fmt.Printf("%s%s\n", scheme, r.Host)
	angagoProxy(w, r, targetHost)
}

// angagoProxy Serve a reverse proxy for a given url
func angagoProxy(w http.ResponseWriter, r *http.Request, target string) {
	remote, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Transport = &customTransport{http.DefaultTransport}
	proxy.ModifyResponse = modifyProxyResponse

	r.URL.Scheme = remote.Scheme
	r.URL.Host = remote.Host
	// r.Host = remote.Host
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(w, r)
}
