package usecase

import (
	"SQ-Assessment/model"
	"SQ-Assessment/model/payload"
	"SQ-Assessment/repository/database"
	"SQ-Assessment/repository/google"
	"SQ-Assessment/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserUsecase interface {
	LoginUser(ctx *gin.Context, Code, State string) (resp payload.LoginUserResponse, err util.Error)
	CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err util.Error)
	GetUser(id uint) (user model.User, err util.Error)
	GetListUsers() (users []model.User, err util.Error)
	UpdateUser(user *model.User) (err util.Error)
	DeleteUser(id uint) (err util.Error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(
	userRepo database.UserRepository,
) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) LoginUser(ctx *gin.Context, Code, State string) (resp payload.LoginUserResponse, err util.Error) {
	tokenRes, errAuth := google.GetGoogleOauthToken(ctx, Code)
	if errAuth != nil {
		err = util.Error{
			Message: errAuth.Error(),
			Err:     errAuth,
			Code:    http.StatusBadGateway,
		}
		return
	}

	googleUser, errGetUser := google.GetGoogleUser(tokenRes.AccessToken)
	if errGetUser != nil {
		err = util.Error{
			Message: errGetUser.Error(),
			Err:     errGetUser,
			Code:    http.StatusBadGateway,
		}
		return
	}

	email := strings.ToLower(googleUser.Email)

	userData := model.User{
		Email:    email,
		Username: googleUser.Name,
	}
	_, err = u.userRepository.GetUserByEmail(email)
	if err.NotNil() {
		fmt.Println("LoginUser: User Does Not Exist")
		if err.Err == gorm.ErrRecordNotFound {
			err = u.userRepository.CreateUser(&userData)
		} else {
			return
		}
	}

	ctx.SetCookie("AccessToken", tokenRes.AccessToken, 60*60, "/", "localhost", false, true)
	resp = payload.LoginUserResponse{
		Email:       userData.Email,
		Username:    userData.Username,
		AccessToken: tokenRes.AccessToken,
	}
	return
}

func (u *userUsecase) CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err util.Error) {
	newUser := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	err = u.userRepository.CreateUser(newUser)
	if err.NotNil() {
		return
	}
	resp = payload.CreateUserResponse{
		UserID: newUser.ID,
	}
	return
}

func (u *userUsecase) GetUser(id uint) (user model.User, err util.Error) {
	user.ID = id
	err = u.userRepository.GetUser(&user)
	if err.NotNil() {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	return
}

func (u *userUsecase) GetListUsers() (users []model.User, err util.Error) {
	users, err = u.userRepository.GetUsers()
	if err.NotNil() {
		fmt.Println("GetListUsers: Error getting users from database")
		return
	}
	return
}

func (u *userUsecase) UpdateUser(user *model.User) (err util.Error) {
	err = u.userRepository.UpdateUser(user)
	if err.NotNil() {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}

func (u *userUsecase) DeleteUser(id uint) (err util.Error) {
	user := model.User{}
	user.ID = id
	err = u.userRepository.DeleteUser(&user)
	if err.NotNil() {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}

	return
}
