package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	UserName    string `json:"username"`
	GroupName   string `json:"group_name"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
	Login       string `json:"login"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	Status      bool   `json:"-"`
	UserIcon    string `json:"-"`
	Action      string `json:"action"`
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan Message
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

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
