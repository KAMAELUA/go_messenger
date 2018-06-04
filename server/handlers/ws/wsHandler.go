package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"../../protocols"
	"github.com/gorilla/websocket"
)


//Client struct ...
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan protocols.Message
}

//ReadOnConnection func ...
func (c *Client) ReadOnConnection() {

	var msg protocols.Message

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		c.hub.broadcast <- msg
	}
}

//WriteOnConnection func ...
func (c *Client) WriteOnConnection() {
	defer func() {
		c.conn.Close()
	}()

	for {
		message, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := c.conn.WriteJSON(message)
		if err != nil {
			fmt.Println("Cannot write json")
		}
	}
}
