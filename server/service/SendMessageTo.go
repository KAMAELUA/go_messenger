package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"../db"
	"../models"
	"../protocols"
)

func DBCONN() *gorm.DB{
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.DB_HOST, db.DB_PORT, db.DB_USER, db.DB_NAME, db.DB_PASSWORD, db.DB_SSLMODE)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
    return db
}

func SendMessageTo(content string, username string, groupname string, contentType string) protocols.Message{
	dbConn := DBCONN()
	sender := models.User{}
	recipient := models.Group{}
	dbConn.Where("username = ?", username).First(&sender)
	dbConn.Where("group_name = ?", groupname).First(&recipient)
	message := models.Message{Content: content,MessageSenderID: sender.ID,MessageRecipientID: recipient.ID,MessageContentType:contentType}
	if dbConn.NewRecord(message){
		dbConn.Create(&message)	
	}
	return protocols.Message{UserName:username,GroupName: recipient.GroupName,ContentType:contentType,Content:content}
}