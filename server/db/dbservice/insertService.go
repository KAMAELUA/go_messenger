package dbservice
	
import (
	"fmt"

	"../../db"
	"../../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBCON struct{
	con *gorm.DB
}

func (dbcon *DBCON) Close(){
	dbcon.con.Close()
}
func NewCon() DBCON{
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.DB_HOST, db.DB_PORT, db.DB_USER, db.DB_NAME, db.DB_PASSWORD, db.DB_SSLMODE)
	temCon, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err)
	}
	return DBCON{
		con: temCon,
	}
}

func (dbcon * DBCON) CreateUser(login string, password string, username string,email string, status bool, usericon string ) bool{
	user := models.User{Login: login, Password: password, Username: username, Email: email, Status: status, UserIcon: usericon}
	if dbcon.con.NewRecord(user){
		dbcon.con.Create(&user)
		return true
	}else{
		return false
	}
}

func (dbcon * DBCON) CreateGroupType(groupType string) bool{
	gtype := models.GroupType{Type: groupType}
	if dbcon.con.NewRecord(gtype){
		dbcon.con.Create(&gtype)
		return true
	}else{
		return false
	}
}
func (dbcon * DBCON) CreateMessageType(messageType string) bool{
	mtype := models.MessageContentType{Type: messageType}
	if dbcon.con.NewRecord(mtype){
		dbcon.con.Create(&mtype)
		return true
	}else{
		return false
	}
}

func (dbcon * DBCON) CreateRelationType(relationType string) bool{
	rtype := models.RelationType{Type: relationType}
	if dbcon.con.NewRecord(rtype){
		dbcon.con.Create(&rtype)
		return true
	}else{
		return false
	}
}

func (dbcon * DBCON) CreateUserRelation(relatingUser string, relatedUser string, relationType uint) bool{
	relatingU := models.User{}
	relatedU := models.User{}
	dbcon.con.Where("username = ?", relatingUser).First(&relatingU)
	dbcon.con.Where("username = ?", relatedUser).First(&relatedU)
	relation := models.UserRelation{RelatingUser:relatingU.ID, RelatedUser: relatedU.ID,RelationTypeID:relationType}
	if dbcon.con.NewRecord(relation){
		dbcon.con.Create(&relation)
		return true
	}else{
		return false
	}
}

func (dbcon * DBCON) CreateGroup(groupName string, groupOwner string, groupType uint) bool{
	owner := models.User{}
	dbcon.con.Where("username = ?", groupOwner).First(&owner)
	group := models.Group{GroupName:groupName,GroupOwnerID: owner.ID,GroupTypeID:groupType}
	if dbcon.con.NewRecord(group){
		dbcon.con.Create(&group)
		return true
	}else{
		return false
	}
}

func (dbcon * DBCON) AddMessage(content string, username string, groupName string, contentType uint) bool{
	sender := models.User{}
	recipient := models.Group{}
	dbcon.con.Where("username = ?", username).First(&sender)
	dbcon.con.Where("group_name = ?", groupName).First(&recipient)
	message := models.Message{Content: content,MessageSenderID: sender.ID,MessageRecipientID: recipient.ID,MessageContentTypeID:contentType}
	if dbcon.con.NewRecord(message){
		dbcon.con.Create(&message)
		return true
	}else{
		return false
	}
}

func (dbcon * DBCON) AddGroupMember(username string, groupName string, lastmessage string) bool{
	user := models.User{}
	group := models.Group{}
	message := models.Message{}
	dbcon.con.Where("username = ?", username).First(&user)
	dbcon.con.Where("group_name = ?", groupName).First(&group)
	dbcon.con.Where("content = ?", lastmessage).First(&message)
	member := models.GroupMember{UserID: user.ID,GroupID: group.ID,LastReadMessageID: message.ID}
	if dbcon.con.NewRecord(member){
		dbcon.con.Create(&member)
		return true
	}else{
		return false
	}

}
