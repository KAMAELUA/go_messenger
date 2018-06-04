package service

import(

)

func SendMessageAll(){
	for conns, name := range userConnections.TCPConnections {
		if(name != message.UserName){
		conns.Write([]byte(desktop_client.JSONencode(message.UserName, "", "", message.Content, "Broadcast", " ", " ", false, " ", "SendMessageTo")))
		}
	}
	for client,_ := range userConnections.WSConnections {
		client.WriteJSON(wsmsg)
	}	
}