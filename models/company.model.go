package models

import (
	"github.com/jinzhu/gorm"
)

type (
	Company struct {
		gorm.Model
		Email       string `json:"email,omitempty" validate:"required,email" gorm:"type:varchar(100);unique_index"`
		Name        string `json:"name,omitempty" validate:"required" gorm:"type:varchar(100)"`
		Address     string `json:"address,omitempty" validate:"required" gorm:"type:varchar(255)"`
		Address2    string `json:"address2,omitempty" gorm:"type:varchar(255)"`
		City        string `json:"city,omitempty" gorm:"type:varchar(255)"`
		State       string `json:"state,omitempty" gorm:"type:varchar(15)"`
		PostalCode  string `json:"postalCode,omitempty" validate:"required" gorm:"type:varchar(10)"`
		PhoneNumber string `json:"phoneNumber,omitempty" validate:"required" gorm:"type:varchar(11)"`
		UserID      uint   `json:"userID,omitempty" validate:"-" gorm:"type:int(10)unsigned"`
	}

	CompanyModel interface {
		GetCompanyById(ID int) (*Company, error)
		GetAllCompanies() ([]Company, error)
		CreateCompany(company Company) (*Company, error)
		UpdateCompany(company Company) (*Company, error)
		DeleteCompany(ID uint) error
	}

	companyModel struct {
		db *gorm.DB
	}
)

func NewCompany(db *gorm.DB) CompanyModel {
	return companyModel{db}
}

func (model companyModel) GetCompanyById(ID int) (*Company, error) {
	var company Company
	err := model.db.First(&company, ID).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (model companyModel) GetAllCompanies() ([]Company, error) {
	var companies []Company
	err := model.db.Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (model companyModel) CreateCompany(company Company) (*Company, error) {
	newCompany := Company{}

	err := model.db.Create(&company).Scan(&newCompany).Error

	if err != nil {
		return nil, err
	}
	return &newCompany, nil
}

func (model companyModel) UpdateCompany(company Company) (*Company, error) {
	updatedCompany := Company{}

	err := model.db.Save(&company).Scan(&updatedCompany).Error

	if err != nil {
		return nil, err
	}
	return &updatedCompany, nil
}

func (model companyModel) DeleteCompany(ID uint) error {
	err := model.db.Unscoped().Delete(Company{Model: gorm.Model{ID: ID}}).Error

	if err != nil {
		return err
	}
	return nil
}
