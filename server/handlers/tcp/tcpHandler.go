package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"../../protocols"
	"../../userConnections"
	"../../routing"
)


var connections []net.Conn

//Handler func ...
func Handler() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP server started on :8080")
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		connections = append(connections, conn)
		if err != nil {
			log.Print("Connection doesn't accepted: ")
			log.Fatal(err)
		}

		go HandleJSON(conn)
	}
}

//HandleJSON func ...
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

//ParseJSON func ..
func ParseJSON(bytes []byte, conn net.Conn) (protocols.Message, string, string) {
	flag := "tcp"
	message := protocols.Message{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	fmt.Println(message.Login)
	fmt.Println(message.Content)
	userConnections.TCPConnections[conn] = message.UserName
	SendToAll(routing.RouterIn(message))
	/*for conns, _ := range userConnections.TCPConnections {
		conns.Write([]byte(JSONencode(message.UserName, "", "", message.Content, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
		conns.Write([]byte("\n"))
	}
	for client,_ := range userConnections.WSConnections {
		client.WriteJSON(message)
	}*/
	return message, "func", flag
}

func SendToAll(msg protocols.Message){
	for conns,_ := range userConnections.TCPConnections {
		conns.Write([]byte(JSONencode(msg.UserName, "", "", msg.Content, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
		conns.Write([]byte("\n"))
	}
	for conns2,_ := range userConnections.WSConnections {
		conns2.WriteJSON(msg)
	}
}