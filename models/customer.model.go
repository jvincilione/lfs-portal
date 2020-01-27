package models

import (
	"github.com/jinzhu/gorm"
)

type (
	Customer struct {
		gorm.Model
		Email       string `json:"email,omitempty" validate:"required,email" gorm:"type:varchar(100);unique_index"`
		FullName    string `json:"fullName,omitempty" validate:"required" gorm:"type:varchar(100)"`
		Address     string `json:"address,omitempty" validate:"required" gorm:"type:varchar(255)"`
		Address2    string `json:"address2,omitempty" gorm:"type:varchar(255)"`
		City        string `json:"city,omitempty" gorm:"type:varchar(255)"`
		State       string `json:"state,omitempty" gorm:"type:varchar(15)"`
		PostalCode  string `json:"postalCode,omitempty" validate:"required" gorm:"type:varchar(10)"`
		PhoneNumber string `json:"phoneNumber,omitempty" validate:"required" gorm:"type:varchar(11)"`
		UserId      uint   `json:"userId,omitempty" validate:"-" gorm:"type:int(10)unsigned"`
	}

	CustomerModel interface {
		GetCustomerById(ID int) (*Customer, error)
		GetAllCustomers() ([]Customer, error)
		CreateCustomer(customer Customer) (*Customer, error)
		UpdateCustomer(customer Customer) (*Customer, error)
		DeleteCustomer(ID uint) error
	}

	customerModel struct {
		db *gorm.DB
	}
)

func NewCustomer(db *gorm.DB) CustomerModel {
	return customerModel{db}
}

func (model customerModel) GetCustomerById(ID int) (*Customer, error) {
	var customer Customer
	err := model.db.First(&customer, ID).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (model customerModel) GetAllCustomers() ([]Customer, error) {
	var customers []Customer
	err := model.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (model customerModel) CreateCustomer(customer Customer) (*Customer, error) {
	newCustomer := Customer{}

	err := model.db.Create(&customer).Scan(&newCustomer).Error

	if err != nil {
		return nil, err
	}
	return &newCustomer, nil
}

func (model customerModel) UpdateCustomer(customer Customer) (*Customer, error) {
	updatedCustomer := Customer{}

	err := model.db.Save(&customer).Scan(&updatedCustomer).Error

	if err != nil {
		return nil, err
	}
	return &updatedCustomer, nil
}

func (model customerModel) DeleteCustomer(ID uint) error {
	err := model.db.Unscoped().Delete(Customer{Model: gorm.Model{ID: ID}}).Error

	if err != nil {
		return err
	}
	return nil
}
