package main

import (
	"./handlers/tcp"
	"flag"
	"log"
	"net/http"
	"./handlers/ws"
)

func tcpHandler(){
	tcp.Handler()
}

func wsHandler(){
	flag.Parse()
	hub := ws.NewHub()

	go hub.RunHub()
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWebsocket(hub, w, r)
	})

	log.Println("HTTP server started on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	go wsHandler()	
	tcpHandler()
}