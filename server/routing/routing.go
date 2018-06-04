package routing

import (
	"fmt"
	"../service"
	"../protocols"
)

func RouterIn(msg protocols.Message) protocols.Message{

	// variable "action" is a command what to do with the structure
	switch msg.Action {

	case "SendMessageTo":
		 return service.SendMessageTo(msg.Content, msg.UserName, msg.GroupName, msg.ContentType)
		 // it doesn't have contentType param
	/*case "CreateUser":
		go service.CreateUser(msg.Login, msg.Password, msg.UserName, msg.Email, msg.UserIcon, msg.Status)
	case "CreateUserRelation":
		CreateUserRelation(msg.RelatingUser, msg.RelatedUser, msg.RelationType) // *1
	case "CreateGroup":
		go service.CreateGroup(msg.GroupName, msg.GroupOwner, msg.GroupMember, msg.GroupType)
	case "AddGroupMember":
		go service.AddGroupMember(msg.UserName, msg.GroupName, msg.LastMessage, msg.GroupMember, msg.GroupType) // why LastMessage is string type?
*/
	default:
		fmt.Println("Unknown format of data")
		return protocols.Message{}
	}
}