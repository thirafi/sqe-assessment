package usecase

import (
	"SQ-Assessment/model"
	"SQ-Assessment/model/payload"
	"SQ-Assessment/util"
	"reflect"
	"testing"

	mockDatabase "SQ-Assessment/mock/repository/database"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_otpUsecase_CreateOtp(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockOtpRepository := mockDatabase.NewMockOtpRepository(mockController)
	mockUserRepository := mockDatabase.NewMockUserRepository(mockController)
	mockOtpUsecase := &otpUsecase{
		otpRepository:  mockOtpRepository,
		userRepository: mockUserRepository,
	}

	returnUser := model.User{
		Model: gorm.Model{
			ID: 123,
		},
		Username: "TestAuthU",
	}

	type args struct {
		username string
	}

	tests := []struct {
		name       string
		otpUsecase *otpUsecase
		args       args
		wantResp   payload.CreateOtpResponse
		wantErr    util.Error
		mock       func()
	}{
		// TODO: Add test cases.
		{
			name:       "Test_otpUsecase_CreateOtp: error GetUserByUsername",
			otpUsecase: mockOtpUsecase,
			args:       args{username: "TestAuthU"},
			wantResp:   payload.CreateOtpResponse{},
			wantErr:    util.Error{Message: "Error to get user by username"},
			mock: func() {
				mockUserRepository.EXPECT().GetUserByUsername(gomock.Any()).Return(model.User{}, util.Error{Message: "Error to get user by username"})
			},
		},
		{
			name:       "Test_otpUsecase_CreateOtp: error GetUserByUsername User not found",
			otpUsecase: mockOtpUsecase,
			args:       args{username: "TestAuthU"},
			wantResp:   payload.CreateOtpResponse{},
			wantErr:    util.Error{Message: "User not found"},
			mock: func() {
				mockUserRepository.EXPECT().GetUserByUsername(gomock.Any()).Return(model.User{}, util.Error{Message: "User not found"})
			},
		},
		{
			name:       "Test_otpUsecase_CreateOtp: error GetActiveOTPByUserId Error update otp",
			otpUsecase: mockOtpUsecase,
			args:       args{username: "TestAuthU"},
			wantResp:   payload.CreateOtpResponse{},
			wantErr:    util.Error{Message: "Error update otp"},
			mock: func() {
				mockUserRepository.EXPECT().GetUserByUsername(gomock.Any()).Return(returnUser, util.Error{})
				mockOtpRepository.EXPECT().GetActiveOTPByUserId(gomock.Any()).Return(model.Otp{}, util.Error{})
				mockOtpRepository.EXPECT().UpdateOTP(gomock.Any()).Return(util.Error{Message: "Error update otp"})
			},
		},
		{
			name:       "Test_otpUsecase_CreateOtp: error GetActiveOTPByUserId Error to create otp",
			otpUsecase: mockOtpUsecase,
			args:       args{username: "TestAuthU"},
			wantResp:   payload.CreateOtpResponse{},
			wantErr:    util.Error{Message: "Error to create otp"},
			mock: func() {
				mockUserRepository.EXPECT().GetUserByUsername(gomock.Any()).Return(returnUser, util.Error{})
				mockOtpRepository.EXPECT().GetActiveOTPByUserId(gomock.Any()).Return(model.Otp{}, util.Error{})
				mockOtpRepository.EXPECT().UpdateOTP(gomock.Any()).Return(util.Error{})
				mockOtpRepository.EXPECT().CreateOTP(gomock.Any()).Return(util.Error{Message: "Error to create otp"})
			},
		},
		{
			name:       "Test_otpUsecase_CreateOtp: SUCCESS GetActiveOTPByUserId still have otp active",
			otpUsecase: mockOtpUsecase,
			args:       args{username: "TestAuthU"},
			wantResp: payload.CreateOtpResponse{
				UserID: "TestAuthU",
			},
			wantErr: util.Error{},
			mock: func() {
				mockUserRepository.EXPECT().GetUserByUsername(gomock.Any()).Return(returnUser, util.Error{})
				mockOtpRepository.EXPECT().GetActiveOTPByUserId(gomock.Any()).Return(model.Otp{}, util.Error{Message: "Error to get otp"})
				mockOtpRepository.EXPECT().CreateOTP(gomock.Any()).Return(util.Error{})
			},
		},
		{
			name:       "Test_otpUsecase_CreateOtp: SUCCESS GetActiveOTPByUserId does`nt have an active otp",
			otpUsecase: mockOtpUsecase,
			args:       args{username: "TestAuthU"},
			wantResp: payload.CreateOtpResponse{
				UserID: "TestAuthU",
			},
			wantErr: util.Error{},
			mock: func() {
				mockUserRepository.EXPECT().GetUserByUsername(gomock.Any()).Return(returnUser, util.Error{})
				mockOtpRepository.EXPECT().GetActiveOTPByUserId(gomock.Any()).Return(model.Otp{}, util.Error{Message: "OTP not found"})
				mockOtpRepository.EXPECT().CreateOTP(gomock.Any()).Return(util.Error{})
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotErr := tt.otpUsecase.CreateOtp(tt.args.username)
			if !reflect.DeepEqual(gotResp.UserID, tt.wantResp.UserID) {
				t.Errorf("otpUsecase.CreateOtp() gotResp = %v, want %v", gotResp.UserID, tt.wantResp.UserID)
			}
			if !reflect.DeepEqual(gotErr.Message, tt.wantErr.Message) {
				t.Errorf("otpUsecase.CreateOtp() gotErr = %v, want %v", gotErr.Message, tt.wantErr.Message)
			}
		})
	}
}
