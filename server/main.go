package main

import (
	"github.com/gokh16/go_messenger/server/handlers/tcp"
	"github.com/gokh16/go_messenger/server/handlers/ws"
)

func main() {
	go ws.RunWebsocket()
	tcp.Handler()

}
