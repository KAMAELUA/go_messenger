package models

import (
	"github.com/jinzhu/gorm"
)

type MessageContentType struct {
	gorm.Model `json:"-"`

	Type string `json:"-"`
}
