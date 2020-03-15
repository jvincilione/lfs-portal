package services

import (
	"fmt"
	"lfs-portal/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		GetUserById(ID int) (*models.PublicUser, error)
		GetUserByEmail(email string) (*models.PublicUser, error)
		GetAllUsers() ([]models.PublicUser, error)
		GetCompanyUsers(companyID uint) ([]models.PublicUser, error)
		CreateUser(user models.User) (*models.PublicUser, error)
		UpdateUser(user models.PublicUser) (*models.PublicUser, error)
		DeleteUser(ID int) error
		AuthenticateUser(email string, password string) string
		UpdatePasswordForUser(user models.AuthUser) error
	}

	userService struct {
		model models.UserModel
	}
)

func NewUserService(model models.UserModel) UserService {
	return userService{model}
}

func (svc userService) GetUserById(ID int) (*models.PublicUser, error) {
	user, err := svc.model.GetUserById(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetUserById] Error getting user, %v", err))
		return nil, err
	}
	return user, nil
}

func (svc userService) GetUserByEmail(email string) (*models.PublicUser, error) {
	user, err := svc.model.GetUserByEmail(email)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetUserByEmail] Error getting user, %v", err))
		return nil, err
	}
	return user, nil
}

func (svc userService) GetAllUsers() ([]models.PublicUser, error) {
	users, err := svc.model.GetAllUsers()
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetAllUsers] Error getting all users, %v", err))
		return nil, err
	}

	return users, nil
}

func (svc userService) GetCompanyUsers(companyID uint) ([]models.PublicUser, error) {
	users, err := svc.model.GetCompanyUsers(companyID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetAllUsers] Error getting all users, %v", err))
		return nil, err
	}

	return users, nil
}

func (svc userService) CreateUser(user models.User) (*models.PublicUser, error) {
	hashedPass := hashAndSalt([]byte(user.Password))
	user.Password = hashedPass
	newUser, err := svc.model.CreateUser(user)
	if err != nil {
		logrus.Error(fmt.Sprintf("[CreateUser] Error creating user, %v", err))
		return nil, err
	}
	return newUser, nil
}

func (svc userService) UpdateUser(user models.PublicUser) (*models.PublicUser, error) {
	updatedUser, err := svc.model.UpdateUser(user)
	if err != nil {
		logrus.Error(fmt.Sprintf("[UpdateUser] Error updating user, %v", err))
		return nil, err
	}
	return updatedUser, nil
}

func (svc userService) UpdatePasswordForUser(user models.AuthUser) error {
	fullUser := models.User{Model: gorm.Model{ID: user.ID}}
	hashedPass := hashAndSalt([]byte(user.Password))
	fullUser.Password = hashedPass
	err := svc.model.UpdateUserPassword(fullUser)
	if err != nil {
		logrus.Error(fmt.Sprintf("[UpdatePasswordForUser] Error updating password for user: %v, Error: %v", user.Email, user.Password))
	}
	return nil
}

func (svc userService) DeleteUser(ID int) error {
	var userID uint
	userID = uint(ID)
	err := svc.model.DeleteUser(userID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[DeleteUser] Error deleting user, %v", err))
		return err
	}
	return nil
}

func (svc userService) AuthenticateUser(email string, password string) string {
	user, err := svc.model.GetUserForAuthentication(email)
	if err != nil {
		logrus.Error(fmt.Sprintf("[AuthenticateUser] Error getting user with email: %v, Error: %v", email, err))
		return ""
	}

	isEqual := comparePasswords(user.Password, []byte(password))
	if !isEqual {
		logrus.Info(fmt.Sprintf("Invalid login attempt, user email: %v", user.Email))
		return ""
	}
	token, err := GenerateJWT(*user)
	return token
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logrus.Error(fmt.Sprintf("[hashAndSalt] Error hashing, %v", err))
		return ""
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logrus.Error(fmt.Sprintf("[comparePasswords] Error comparing password to hashed password, %v", err))
		return false
	}

	return true
}
