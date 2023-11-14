package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var resp []byte
		if req.URL.Path == "/handler-initial-data" {
			resp = []byte(`{"text": "initial"}`)
		} else if req.URL.Path == "/handler" {
			time.Sleep((time.Second)) // todo hack sleep a second to check if it's ok
			resp = []byte(`{"text": "updated"}`)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		} else if req.URL.Path == "/username" {
			resp = []byte(`{"username": "colin"}`)
		} else {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
		rw.Write(resp)
	})

	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
