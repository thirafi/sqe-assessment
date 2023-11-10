package database

import (
	"SQ-Assessment/config"
	"SQ-Assessment/model"
	"SQ-Assessment/util"

	"gorm.io/gorm"
)

type OtpRepository interface {
	CreateOTP(otp *model.Otp) util.Error
	GetOTPByUserId(userId uint, otpNumber string) (otp model.Otp, err util.Error)
	GetActiveOTPByUserId(userId uint) (otp model.Otp, err util.Error)
	UpdateOTP(otp *model.Otp) util.Error
}

type otpRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) *otpRepository {
	return &otpRepository{db}
}

func (u *otpRepository) CreateOTP(otp *model.Otp) (err util.Error) {
	if errDb := config.DB.Create(otp).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to create otp")
		return
	}
	return
}

func (u *otpRepository) GetOTPByUserId(userId uint, otpNumber string) (otp model.Otp, err util.Error) {
	if errDb := config.DB.Where("user_id = ? AND otp = ? AND status = 'created' ", userId, otpNumber).First(&otp).Error; errDb != nil {
		if errDb == gorm.ErrRecordNotFound {
			err.StatusNotFound(errDb, "OTP not found")
			return
		}
		err.StatusInternalServerError(errDb, "Error to get otp")
		return
	}
	return
}

func (u *otpRepository) GetActiveOTPByUserId(userId uint) (otp model.Otp, err util.Error) {
	if errDb := config.DB.Where("user_id = ? AND status = 'created' ", userId).First(&otp).Error; errDb != nil {
		if errDb == gorm.ErrRecordNotFound {
			err.StatusNotFound(errDb, "OTP not found")
			return
		}
		err.StatusInternalServerError(errDb, "Error to get otp")
		return
	}
	return
}

func (u *otpRepository) UpdateOTP(otp *model.Otp) (err util.Error) {
	if errDb := config.DB.Updates(otp).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error update otp")
		return err
	}
	return
}
