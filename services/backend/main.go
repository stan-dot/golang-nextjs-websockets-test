package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
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
		} else if req.URL.Path == "/socket" {
			conn, _, _, err := ws.UpgradeHTTP(req, rw)
			if err != nil {
				log.Println("Error with WebSocket: ", err)
				rw.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			go func() {
				defer conn.Close()
				time.Sleep(time.Second)
				err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(`{"text": "from-websocket"}`))
				if err != nil {
					log.Println("error writing WebSocket data: ", err)
					return
				}
			}()
			return
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
