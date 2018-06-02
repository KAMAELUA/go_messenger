package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"-"`

	Login    string
	Password string `json:"-"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   bool   `json:"-"`
	UserIcon string `json:"-"`
}
