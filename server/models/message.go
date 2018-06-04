package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model `json:"-"`

	User               User
	Group              Group

	Content              string `json:"message_content"`
	MessageSenderID      uint   `json:"-"`
	MessageRecipientID   uint   `json:"-"`
	MessageContentType	 string   `json:"-"`
}
