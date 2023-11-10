package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"size:255;unique"`
	Email    string `json:"email" form:"email" gorm:"size:255;unique"`
	Password string `json:"-" form:"password" gorm:"size:255"`
}
