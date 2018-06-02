package models

import (
	"github.com/jinzhu/gorm"
)

type RelationType struct {
	gorm.Model `json:"-"`

	Type string `json:"-"`
}
