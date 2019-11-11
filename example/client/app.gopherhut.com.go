package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	port := ":8081"
	http.HandleFunc("/", handler)
	log.Println("app.gopherhut.com listening on ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host + r.URL.Path)
	res := map[string]interface{}{
		"host": r.Host,
		"data": map[string]interface{}{
			"incident": "12",
			"message":  "everything fuckedup!",
		},
		"meta": map[string]interface{}{
			"status":      "ok",
			"status_code": http.StatusOK,
		},
	}

	resp, err := json.Marshal(&res)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	// w.Header().Set("Host", r.Host)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Fatal(err.Error())
	}
}
