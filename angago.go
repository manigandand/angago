package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func angagoServer(w http.ResponseWriter, r *http.Request) {
	targetURL, err := angagoConfig.TargetURL(r.Host)
	if err != nil {
		if err := respondFail(w, r, err); err != nil {
			log.Println(err.Error())
		}
		return
	}

	log.Printf("Proxying the Host %s to %s\n", r.Host, targetURL)
	angagoProxy(w, r, angagoConfig.AbsTargetURL(targetURL))
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

func respondFail(w http.ResponseWriter, r *http.Request, err error) error {
	res := map[string]interface{}{
		"host": r.Host,
		"meta": map[string]interface{}{
			"status":      http.StatusText(http.StatusNotFound),
			"status_code": http.StatusNotFound,
			"error":       err.Error(),
		},
	}

	resp, err := json.Marshal(&res)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNotFound)
	// w.Header().Set("Host", r.Host)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		return err
	}
	return nil
}
