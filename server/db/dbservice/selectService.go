package dbservice
	
import (
	"../../models"
)

func (dbcon * DBCON) GetMessages(count int, groupName string) []models.Message{
	group := models.Group{}
	messages := []models.Message{}
	dbcon.con.Where("group_name = ?", groupName).First(&group)
	dbcon.con.Where("message_recipient_id = ?", group.ID).Limit(count).Find(&messages)
	return messages
}

func (dbcon * DBCON) GetGroupList(username string) []models.Group{
	user := models.User{}
	dbcon.con.Where("username = ?", username).First(&user)
	groups := []models.Group{}
	dbcon.con.Joins("join group_members on groups.id=group_members.group_id").Where("user_id = ?", user.ID).Find(&groups) 
	return groups
}

func (dbcon * DBCON)GetGroupUserList(groupName string) []models.User{
	group := models.Group{}
	users := []models.User{}
	dbcon.con.Where("group_name = ?", groupName).First(&group)
	dbcon.con.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&users)
	return users
}
func (dbcon * DBCON) FindUser(username string) models.User{
	user := models.User{}
	dbcon.con.Where("username = ?", username).First(&user)
	return user
}

func (dbcon * DBCON) GetContactList(username string) []models.User{
	user := models.User{}
	temp := []models.UserRelation{}
	friends := []models.User{}
	dbcon.con.Where("username = ?", username).First(&user)
	dbcon.con.Where("relating_user=?",user.ID).Find(&temp)
	for i,_:= range temp{
		friend := models.User{}
		dbcon.con.Where("id=?",temp[i].RelatedUser).First(&friend)
		friends = append(friends, friend)
		
	}
	return friends

}
