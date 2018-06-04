package main

import (
	"./handlers/tcp"
	"./handlers/ws"
	"./db"
)
func init(){
	db.CreateDatabase()
}
func main() {
	go ws.RunWebsocket()
	tcp.Handler()

}
