package models

import (
	"github.com/jinzhu/gorm"
)

type GroupType struct {
	gorm.Model `json:"-"`

	Type string `json:"-"`
}
