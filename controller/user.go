package controller

import (
	"SQ-Assessment/model"
	"SQ-Assessment/model/payload"
	"SQ-Assessment/repository/database"
	"SQ-Assessment/usecase"
	"SQ-Assessment/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsercontroller(c *gin.Context) error
	GetUserController(c *gin.Context) error
	CreateUserController(c *gin.Context) error
	DeleteUserController(c *gin.Context) error
	UpdateUserController(c *gin.Context) error
}

type userController struct {
	userUsecase    usecase.UserUsecase
	userRepository database.UserRepository
	otpUsecase     usecase.OtpUsecase
	otpRepository  database.OtpRepository
}

func NewUserController(
	userUsecase usecase.UserUsecase,
	userRepository database.UserRepository,
	otpUsecase usecase.OtpUsecase,
	otpRepository database.OtpRepository,
) *userController {
	return &userController{
		userUsecase,
		userRepository,
		otpUsecase,
		otpRepository,
	}
}

func (u *userController) GetUsersController(c *gin.Context) {
	users, err := u.userUsecase.GetListUsers()
	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func (u *userController) GetAllUsersController(c *gin.Context) {
	user, err := u.userUsecase.GetListUsers()

	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})

}

func (u *userController) GetUserController(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := u.userUsecase.GetUser(uint(id))

	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})

}

// create new user
func (u *userController) CreateUserController(c *gin.Context) {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	// validasi request body
	if errBinding := c.ShouldBind(payload); errBinding != nil {
		c.JSON(http.StatusBadRequest, util.Error{
			Code:    http.StatusBadRequest,
			Err:     errBinding,
			Message: "invalid payload",
		})
		return
	}
	resp, err := u.userUsecase.CreateUser(&payload)
	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    resp,
	})
}

// delete user by id
func (u *userController) DeleteUserController(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := u.userUsecase.DeleteUser(uint(id)); err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})

}

// update user by id
func (u *userController) UpdateUserController(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user := model.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := u.userUsecase.UpdateUser(&user); err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
