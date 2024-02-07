package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	var resp = &Response{Message: "hello world!"}
	j, _ := json.Marshal(resp)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write(j)
	}

	healthHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok!"))
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
