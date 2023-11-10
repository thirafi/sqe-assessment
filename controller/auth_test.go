package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockDatabase "SQ-Assessment/mock/repository/database"
	mockUsecase "SQ-Assessment/mock/usecase"
	"SQ-Assessment/model/payload"
	"SQ-Assessment/util"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func generateContext(method, url string, body []byte) (ctx *gin.Context) {
	gin.SetMode(gin.TestMode)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(body))
		req.Header.Add("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req
	return
}

func Test_userController_ActionOtpController(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockUserUsecase := mockUsecase.NewMockUserUsecase(mockController)
	mockOtpUsecase := mockUsecase.NewMockOtpUsecase(mockController)
	mockOtpRepository := mockDatabase.NewMockOtpRepository(mockController)
	mockUserRepository := mockDatabase.NewMockUserRepository(mockController)
	mockUserController := &userController{
		userUsecase:    mockUserUsecase,
		userRepository: mockUserRepository,
		otpUsecase:     mockOtpUsecase,
		otpRepository:  mockOtpRepository,
	}

	method := http.MethodPost
	url := "/api/otp"
	payloadCreateOtp, _ := json.Marshal(payload.OtpActionRequest{Username: "test1", Action: "create"})
	payloadValidateOtp, _ := json.Marshal(payload.OtpActionRequest{Username: "test1", Action: "validate"})
	payloadInvalidAction, _ := json.Marshal(payload.OtpActionRequest{Username: "test2", Action: "insert"})
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name           string
		userController *userController
		args           args
		mock           func()
	}{
		// TODO: Add test cases.
		{
			name:           "Test_userController_ActionOtpController: error bind json payload",
			userController: mockUserController,
			args: args{
				c: generateContext(method, url, nil),
			},
			mock: func() {},
		},
		{
			name:           "Test_userController_ActionOtpController: error Create OTP",
			userController: mockUserController,
			args: args{
				c: generateContext(method, url, payloadCreateOtp),
			},
			mock: func() {
				mockOtpUsecase.EXPECT().CreateOtp(gomock.Any()).Return(payload.CreateOtpResponse{}, util.Error{Message: "Error creating user"})
			},
		},
		{
			name:           "Test_userController_ActionOtpController: Success Create OTP",
			userController: mockUserController,
			args: args{
				c: generateContext(method, url, payloadCreateOtp),
			},
			mock: func() {
				mockOtpUsecase.EXPECT().CreateOtp(gomock.Any()).Return(payload.CreateOtpResponse{UserID: "test1", OTP: "521321"}, util.Error{})
			},
		},
		{
			name:           "Test_userController_ActionOtpController: error Validate OTP",
			userController: mockUserController,
			args: args{
				c: generateContext(method, url, payloadValidateOtp),
			},
			mock: func() {
				mockOtpUsecase.EXPECT().ValidateOtp(gomock.Any()).Return(payload.ValidateOtpResponse{}, util.Error{Message: "Error to get user by username"})
			},
		},
		{
			name:           "Test_userController_ActionOtpController: Success Validate OTP",
			userController: mockUserController,
			args: args{
				c: generateContext(method, url, payloadValidateOtp),
			},
			mock: func() {
				mockOtpUsecase.EXPECT().ValidateOtp(gomock.Any()).Return(payload.ValidateOtpResponse{UserID: "test1", Message: "OTP validated successfully"}, util.Error{})
			},
		},
		{
			name:           "Test_userController_ActionOtpController: Error Action is invalid",
			userController: mockUserController,
			args: args{
				c: generateContext(method, url, payloadInvalidAction),
			},
			mock: func() {},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tt.userController.ActionOtpController(tt.args.c)
		})
	}
}
