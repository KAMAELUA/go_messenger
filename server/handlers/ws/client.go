package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gokh16/go_messenger/server/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan models.Message
}

func (c *Client) ReadOnConnection() {

	var msg models.Message

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
			fmt.Println("There was an error:", err)
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

func ServeWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Cannot upgrade")
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan models.Message),
	}
	client.hub.register <- client

	go client.WriteOnConnection()
	go client.ReadOnConnection()
}
