package ws

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//Message struct ...
type Message struct {
	UserName    string `json:"username"`
	GroupName   string `json:"group_name"`
	ContentType string `json:"content_type"`
	Content     string `json:"message_content"`
	Login       string `json:"login"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	Status      bool   `json:"-"`
	UserIcon    string `json:"-"`
	Action      string `json:"action"`
}

//Client struct ...
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan Message
}

//ReadOnConnection func ...
func (c *Client) ReadOnConnection() {

	var msg Message

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
