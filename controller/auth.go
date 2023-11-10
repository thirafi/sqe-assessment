package controller

import (
	"errors"
	"net/http"

	"SQ-Assessment/model/payload"
	"SQ-Assessment/usecase"
	"SQ-Assessment/util"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginUserController(c *gin.Context) error
}

type authController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(userUsecase usecase.UserUsecase) *authController {
	return &authController{
		userUsecase,
	}
}

func (a *authController) LoginUserController(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	userData, err := a.userUsecase.LoginUser(c, code, pathUrl)
	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   userData,
	})

}

// create new user
func (u *userController) ActionOtpController(c *gin.Context) {
	payload := payload.OtpActionRequest{}
	// validasi request body
	if errBinding := c.ShouldBindJSON(&payload); errBinding != nil {
		c.JSON(http.StatusBadRequest, util.Error{
			Code:    http.StatusBadRequest,
			Err:     errBinding,
			Message: "invalid payload",
		})
		return
	}
	switch payload.Action {
	case "create":
		resp, err := u.otpUsecase.CreateOtp(payload.Username)
		if err.NotNil() {
			c.JSON(err.Code, err)
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success generate otp request",
			"data":    resp,
		})
	case "validate":
		resp, err := u.otpUsecase.ValidateOtp(payload)
		if err.NotNil() {
			c.JSON(err.Code, err)
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OTP validated successfully",
			"data":    resp,
		})
	default:
		c.JSON(http.StatusBadRequest, util.Error{
			Code:    http.StatusBadRequest,
			Err:     errors.New("invalid action"),
			Message: "invalid action",
		})
		return
	}

}
