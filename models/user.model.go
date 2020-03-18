package models

import (
	"lfs-portal/enums"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type (
	User struct {
		gorm.Model
		FirstName string         `json:"firstName,omitempty" validate:"required" gorm:"type:varchar(100);not null"`
		LastName  string         `json:"lastName,omitempty" validate:"required" gorm:"type:varchar(100);not null"`
		Email     string         `json:"email,omitempty" validate:"required,email" gorm:"type:varchar(100);unique_index"`
		Phone     string         `json:"phone,omitempty" validate:"-" gorm:"type:varchar(15)"`
		Title     string         `json:"title,omitempty" validate:"-" gorm:"type:varchar(255)"`
		UserType  enums.UserType `json:"userType,omitempty" gorm:"type:int;not null"`
		Password  string         `json:"password,omitempty" validate:"required" gorm:"type:text"`
		Company   []Company      `json:"companies" gorm:"ForeignKey:UserID"`
		CompanyID uint           `json:"companyId" gorm:"type:int"`
	}

	PublicUser struct {
		ID        uint
		CreatedAt time.Time
		UpdatedAt time.Time
		FirstName string         `json:"firstName,omitempty"`
		LastName  string         `json:"lastName,omitempty"`
		Phone     string         `json:"phone,omitempty"`
		Email     string         `json:"email,omitempty"`
		Title     string         `json:"title,omitempty"`
		UserType  enums.UserType `json:"userType"`
		Company   []Company      `json:"companies" gorm:"ForeignKey:UserID"`
		CompanyID uint           `json:"companyId" gorm:"type:int"`
	}

	AuthUser struct {
		ID        uint
		Email     string         `json:"email" validate:"required"`
		Password  string         `json:"password" validate:"required"`
		FirstName string         `json:"firstName" validate:"-"`
		LastName  string         `json:"lastName" validate:"-"`
		UserType  enums.UserType `json:"userType,omitempty"`
		Companies []uint         `json:"companies,omitempty"`
		CompanyID uint           `json:"companyId" gorm:"type:int"`
	}

	UserModel interface {
		GetUserById(ID int) (*PublicUser, error)
		GetUserByEmail(email string) (*PublicUser, error)
		GetUserForAuthentication(email string) (*AuthUser, error)
		GetAllUsers() ([]PublicUser, error)
		GetCompanyUsers(companyID uint) ([]PublicUser, error)
		CreateUser(user User) (*PublicUser, error)
		UpdateUser(user PublicUser) (*PublicUser, error)
		UpdateUserPassword(user User) error
		DeleteUser(ID uint) error
	}

	userModel struct {
		db *gorm.DB
	}
)

func NewUser(db *gorm.DB) UserModel {
	return userModel{db}
}

func (PublicUser) TableName() string {
	return "users"
}

func (model userModel) GetUserById(ID int) (*PublicUser, error) {
	var user PublicUser
	err := model.db.Table("users").First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (model userModel) GetUserByEmail(email string) (*PublicUser, error) {
	var user PublicUser
	err := model.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (model userModel) GetUserForAuthentication(email string) (*AuthUser, error) {
	var user AuthUser
	err := model.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	var companies []Company
	err = model.db.Table("companies").Where("user_id = ?", user.ID).Find(&companies).Error
	if err != nil {
		return nil, err
	}
	companyIDs := []uint{}
	for _, company := range companies {
		companyIDs = append(companyIDs, company.ID)
	}
	companyIDs = append(companyIDs, user.CompanyID)
	user.Companies = companyIDs
	logrus.Errorf("user id %v", user.ID)
	return &user, nil
}

func (model userModel) GetAllUsers() ([]PublicUser, error) {
	var userList []PublicUser
	err := model.db.Preload("Company").Find(&userList).Error

	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (model userModel) GetCompanyUsers(companyID uint) ([]PublicUser, error) {
	var userList []PublicUser
	err := model.db.Preload("Company").Where("company_id = ?", companyID).Find(&userList).Error

	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (model userModel) CreateUser(user User) (*PublicUser, error) {
	var newUser PublicUser

	err := model.db.Create(&user).Scan(&newUser).Error

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (model userModel) UpdateUser(user PublicUser) (*PublicUser, error) {
	updatedUser := PublicUser{}

	err := model.db.Table("users").Save(&user).Scan(&updatedUser).Error

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (model userModel) UpdateUserPassword(user User) error {
	err := model.db.Table("users").Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (model userModel) DeleteUser(ID uint) error {
	err := model.db.Unscoped().Delete(User{Model: gorm.Model{ID: ID}}).Error

	if err != nil {
		return err
	}
	return nil
}
