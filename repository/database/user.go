package database

import (
	"SQ-Assessment/model"
	"SQ-Assessment/util"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) (err util.Error)
	GetUsers() (users []model.User, err util.Error)
	GetUser(user *model.User) (err util.Error)
	GetUserByEmail(email string) (user model.User, err util.Error)
	GetUserByUsername(username string) (user model.User, err util.Error)
	UpdateUser(user *model.User) (err util.Error)
	DeleteUser(user *model.User) (err util.Error)
	LoginUser(user *model.User) (err util.Error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *model.User) (err util.Error) {
	if errDb := u.db.Create(user).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error creating user")
		return
	}
	return
}

func (u *userRepository) GetUsers() (users []model.User, err util.Error) {
	if errDb := u.db.Model(&model.User{}).Find(&users).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to create user")
		return
	}
	return
}

func (u *userRepository) GetUser(user *model.User) (err util.Error) {
	if errDb := u.db.First(&user).Error; errDb != nil {
		if errDb == gorm.ErrRecordNotFound {
			err.StatusNotFound(errDb, "User not found")
			return
		}
		err.StatusInternalServerError(errDb, "Error to get user")
		return
	}
	return
}

func (u *userRepository) UpdateUser(user *model.User) (err util.Error) {
	if errDb := u.db.Updates(user).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to update user")
		return err
	}
	return
}

func (u *userRepository) DeleteUser(user *model.User) (err util.Error) {
	if errDb := u.db.Delete(user).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to delete user")
		return err
	}
	return
}

func (u *userRepository) LoginUser(user *model.User) (err util.Error) {
	if errDb := u.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error login user")
		return err
	}
	return
}

func (u *userRepository) GetUserByEmail(email string) (user model.User, err util.Error) {
	if errDb := u.db.Where("email = ?", email).First(&user).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to get user by email")
		return
	}
	return
}

func (u *userRepository) GetUserByUsername(username string) (user model.User, err util.Error) {
	if errDb := u.db.Where("username = ?", username).First(&user).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to get user by username")
		return
	}
	return
}
