package model

import (
	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	OTP    string `json:"otp" form:"otp" gorm:"size:50"`
	Status string `json:"status" form:"status" gorm:"type:enum('created', 'validated', 'expired')"`
	UserID uint   `json:"user_id" form:"user_id"`
	User   User
}
