package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	port := ":8081"
	http.HandleFunc("/", handler)
	fmt.Println("app.gopherhut.com listening on ", port)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Host)
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
		fmt.Println("error:", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
