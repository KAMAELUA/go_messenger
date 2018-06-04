package userConnections

import (
	"net"

	"github.com/gorilla/websocket"
)

// WSConnections login
var WSConnections = map[*websocket.Conn]string{}

// TCPConnections login
var TCPConnections = map[net.Conn]string{}
