package route

import (
	"SQ-Assessment/controller"
	"SQ-Assessment/repository/database"
	"SQ-Assessment/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(e *gin.Engine, db *gorm.DB) {
	// e.Validator = &CustomValidator{validator: validator.New()}
	userRepository := database.NewUserRepository(db)
	otpRepository := database.NewOTPRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	otpUsecase := usecase.NewOtpUsecase(otpRepository, userRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository, otpUsecase, otpRepository)

	// middleware := middleware.NewMiddleware(userRepository)

	e.POST("/login", authController.LoginUserController)
	e.GET("/callback", authController.LoginUserController)
	e.POST("/register", userController.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", userController.GetAllUsersController)
	user.GET("/:id", userController.GetUserController)
	user.POST("", userController.CreateUserController)
	user.PUT("/:id", userController.UpdateUserController)
	user.DELETE("/:id", userController.DeleteUserController)

	// otp collection
	otp := e.Group("/otp")
	otp.POST("", userController.ActionOtpController)
}
