package models

import (
	"github.com/jinzhu/gorm"
)

type UserRelation struct {
	gorm.Model `json:"-"`

	RelationType RelationType

	RelatingUser   uint `sql:"type:int REFERENCES users(id)"`
	RelatedUser    uint `sql:"type:int REFERENCES users(id)"`
	RelationTypeID uint
}
