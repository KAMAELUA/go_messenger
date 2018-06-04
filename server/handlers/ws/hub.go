package ws

import (
	"github.com/gokh16/go_messenger/server/userConnections"
)

//Hub struct ...
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

//NewHub func ...
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//RunHub func ..
func (h *Hub) runHub() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			userConnections.WSConnections[client.conn] = ""
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			for conns := range userConnections.TCPConnections {
				conns.Write([]byte(message.Content))
				conns.Write([]byte("\n"))
			}
		}
	}
}
