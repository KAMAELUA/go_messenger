package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"../../userConnections"
	"log"
	"net"
	"../../models"
	"../../../desktop/desktop-client"
)

type Message struct {
	UserName    string
	GroupName   string
	ContentType string
	Content     string
	Login       string
	Password    string
	Email       string
	Status      bool
	UserIcon    string
	Action      string
}

var connections []net.Conn

func Handler() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		log.Print("TCP server started on :8080")
		conn, err := ln.Accept()
		connections = append(connections, conn)

		if err != nil {
			log.Print("Connection doesn't accepted: ")
			log.Fatal(err)
		}

		go HandleJSON(conn)
	}
}

func HandleJSON(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		log.Println(data, err)
		if err != nil {
			log.Print("Data didn't read right: ")
			log.Fatal(err)
		}
		ParseJSON([]byte(data), conn)
	}
}

func ParseJSON(bytes []byte, conn net.Conn) (Message, string) {
	message := Message{}
	wsmsg := models.Message{}
	err := json.Unmarshal(bytes, &message)
	wsmsg.Content = message.Content
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	fmt.Println(message.UserName)
	fmt.Println(message.Content)
	userConnections.TCPConnections[conn] = message.UserName
	for conns, name := range userConnections.TCPConnections {
		if(name != message.UserName){
		conns.Write([]byte(desktop_client.JSONencode(message.UserName, "", "", message.Content, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
		}
	}
	for client,_ := range userConnections.WSConnections {
		client.WriteJSON(wsmsg)
	}		
	return message, " func "
}