package main

import (
	"./handlers/tcp"
	"./handlers/ws"
)

func main() {
	go ws.RunWebsocket()
	tcp.Handler()

}
