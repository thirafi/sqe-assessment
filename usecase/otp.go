package usecase

import (
	"SQ-Assessment/constant"
	"SQ-Assessment/model"
	"SQ-Assessment/model/payload"
	"SQ-Assessment/repository/database"
	"SQ-Assessment/util"
	"errors"
	"fmt"
	"time"
)

type OtpUsecase interface {
	CreateOtp(username string) (resp payload.CreateOtpResponse, err util.Error)
	UpdateOtp(otp *model.Otp) (err util.Error)
	ValidateOtp(req payload.OtpActionRequest) (resp payload.ValidateOtpResponse, err util.Error)
}

type otpUsecase struct {
	otpRepository  database.OtpRepository
	userRepository database.UserRepository
}

func NewOtpUsecase(
	otpRepository database.OtpRepository,
	userRepository database.UserRepository,
) *otpUsecase {
	return &otpUsecase{
		otpRepository:  otpRepository,
		userRepository: userRepository,
	}
}

func (b *otpUsecase) CreateOtp(username string) (resp payload.CreateOtpResponse, err util.Error) {

	// get user data from user repo by username
	userData, err := b.userRepository.GetUserByUsername(username)
	if err.NotNil() {
		fmt.Println("CreateOtp: error get user from database, err:", err)
		return
	}

	activeOtp, err := b.otpRepository.GetActiveOTPByUserId(userData.ID)
	if !err.NotNil() {
		activeOtp.Status = constant.OtpStatusExpired
		if err = b.otpRepository.UpdateOTP(&activeOtp); err.NotNil() {
			fmt.Println("CrateOtp: error create otp to database, err:", err)
			return
		}
	}

	otp := util.EncodeToString(6)
	// save otp to db
	if err = b.otpRepository.CreateOTP(&model.Otp{
		UserID: userData.ID,
		Status: "created",
		OTP:    otp, // generate encode 6 int for otp
	}); err.NotNil() {
		fmt.Println("CreateOtp: error create otp to database, err:", err)
		return
	}
	resp.UserID = userData.Username
	resp.OTP = otp

	return
}

func (b *otpUsecase) ValidateOtp(req payload.OtpActionRequest) (resp payload.ValidateOtpResponse, err util.Error) {

	// get user data from user repo by username
	userData, err := b.userRepository.GetUserByUsername(req.Username)
	if err.NotNil() {
		fmt.Println("ValidateOtp: error get user from database, err:", err)
		return
	}

	otpData, err := b.otpRepository.GetOTPByUserId(userData.ID, req.OTP)
	if err.NotNil() {
		fmt.Println("ValidateOtp: error get otp from database, err:", err)
		return
	}

	// validate expireTime
	now := time.Now().UTC()
	otpData.Status = constant.OtpStatusValidated
	if now.After(otpData.CreatedAt.Add(time.Minute * 2)) {
		fmt.Println("CrateOtp: otp expired")
		err.StatusBadRequest(errors.New("otp expired"), "otp could not be processed")
		otpData.Status = constant.OtpStatusExpired
	}

	// save otp to db
	if err = b.otpRepository.UpdateOTP(&otpData); err.NotNil() {
		fmt.Println("CrateOtp: error create otp to database, err:", err)
		return
	}

	resp.UserID = userData.Username
	resp.Message = "OTP validated successfully"

	return
}

func (b *otpUsecase) UpdateOtp(otp *model.Otp) (err util.Error) {
	err = b.otpRepository.UpdateOTP(otp)
	if err.NotNil() {
		fmt.Println("UpdateOtp : Error updating otp, err: ", err)
		return
	}

	return
}
