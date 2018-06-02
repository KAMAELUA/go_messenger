package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model `json:"-"`

	User               User
	Group              Group
	MessageContentType MessageContentType

	Content              string `json:"message_content"`
	MessageSenderID      uint   `json:"-"`
	MessageRecipientID   uint   `json:"-"`
	MessageContentTypeID uint   `json:"-"`
}
