package payload

type (
	CreateUserRequest struct {
		Username string `json:"username" form:"username" binding:"required,max=20"`
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required,min=5"`
	}

	CreateUserResponse struct {
		UserID uint   `json:"user_id"`
		Token  string `json:"token"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required,min=5"`
	}

	LoginUserResponse struct {
		Email       string `json:"email"`
		Username    string `json:"username"`
		AccessToken string `json:"access_token"`
	}

	OtpActionRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Action   string `json:"action" form:"action" binding:"required"`
		OTP      string `json:"otp" form:"otp"`
	}

	CreateOtpResponse struct {
		UserID string `json:"user_id" form:"user_id"`
		OTP    string `json:"otp" form:"otp"`
	}

	ValidateOtpResponse struct {
		UserID  string `json:"user_id" form:"user_id"`
		Message string `json:"message" form:"message"`
	}
)
