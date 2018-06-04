package connections

import (
	"net"

	"github.com/gorilla/websocket"
)

var WSConnections = map[*websocket.Conn]string{} // connection:login
var TCPConnections = map[net.Conn]string{}       // connection:login
