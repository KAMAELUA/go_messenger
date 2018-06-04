package tcp

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"../../protocols"
)

func JSONencode(user string, groupName string, contentType string, content string, login string, password string, email string, status bool, userIcon string, action string) string {
	incomingData := protocols.Message{user, groupName, contentType, content, login, password, email, status, userIcon, action}
	outcomingData, err := json.Marshal(incomingData)
	if err != nil {
		log.Fatal(err)
	}
	return string(outcomingData) + "\n"
}

func JSONdecode(conn net.Conn) protocols.Message {
	message := protocols.Message{}
	jsonObj, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonObj, &message)
	return message
}