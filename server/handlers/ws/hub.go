package ws

import (
	"../../protocols"
	"../../userConnections"
	"../../routing"
	"../tcp"
)

//Hub struct ...
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan protocols.Message
	register   chan *Client
	unregister chan *Client
}

//NewHub func ...
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan protocols.Message),
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
			userConnections.WSConnections[client.conn] = "UserName"
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			//userConnections.WSConnections[] = message.UserName
			message.Action="SendMessageTo"
			SendToAll(routing.RouterIn(message))
		/*	for client := range h.clients {
				select {
				case client.send <- message:
					for conns,_ := range userConnections.TCPConnections {
						conns.Write([]byte(tcp.JSONencode(message.UserName, "", "", message.Content, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
						conns.Write([]byte("\n"))
					}
					userConnections.WSConnections[client.conn] = message.UserName
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}*/
			
		}
	}
}

func SendToAll(msg protocols.Message){
	for conns,_ := range userConnections.TCPConnections {
		conns.Write([]byte(tcp.JSONencode(msg.UserName, "", "", msg.Content, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
		conns.Write([]byte("\n"))
	}
	for conns,_ := range userConnections.WSConnections {
		conns.WriteJSON(msg)
	}
}