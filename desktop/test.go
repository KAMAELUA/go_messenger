package main


import (
	"bufio"
	"fmt"
	"net"
	
	"os"
	"./desktop-client"
)

const(
	tcpProtocol	= "tcp"
	keySize = 1024
	readWriterSize = keySize/8
)

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var connectAddr = &net.TCPAddr{IP: net.IPv4(0,0,0,0), Port: 8080}

func getMessage(c *net.TCPConn){
	for{
		msg := desktop_client.JSONdecode(c)
		fmt.Println(msg.UserName," ", msg.Content)
	}
}


func connectTo() *net.TCPConn{
	fmt.Print("Enter port:")
	fmt.Scanf("%d", &connectAddr.Port)
	fmt.Println("Connect to", connectAddr)
	c ,err := net.DialTCP(tcpProtocol, nil, connectAddr); checkErr(err)
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ваше имя? ")
	UName, _ := reader.ReadString('\n')
	conn := connectTo()

	fmt.Println("Сообщения: ")
	var message string
	var tmp string
	go getMessage(conn)
	for {
		tmp = " "
		fmt.Scanf("%s",&tmp)
		message = tmp
		if(message != "" && message !=" "){
		conn.Write([]byte(desktop_client.JSONencode(UName, "", "", message, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
		message =" "
		}

	}
}