package services

import (
	"fmt"
	"lfs-portal/models"

	"github.com/sirupsen/logrus"
)

type (
	CustomerService interface {
		GetCustomerById(ID int) (*models.Customer, error)
		GetAllCustomers() ([]models.Customer, error)
		CreateCustomer(customer models.Customer) (*models.Customer, error)
		UpdateCustomer(customer models.Customer) (*models.Customer, error)
		DeleteCustomer(ID int) error
	}

	customerService struct {
		model models.CustomerModel
	}
)

func NewCustomerService(model models.CustomerModel) CustomerService {
	return customerService{model}
}

func (svc customerService) GetCustomerById(ID int) (*models.Customer, error) {
	customer, err := svc.model.GetCustomerById(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetCustomerById] Error getting customer, %v", err))
		return nil, err
	}
	return customer, nil
}

func (svc customerService) GetAllCustomers() ([]models.Customer, error) {
	customers, err := svc.model.GetAllCustomers()
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetAllCustomers] Error getting customers, %v", err))
		return nil, err
	}
	return customers, nil
}

func (svc customerService) CreateCustomer(customer models.Customer) (*models.Customer, error) {
	newCustomer, err := svc.model.CreateCustomer(customer)
	if err != nil {
		logrus.Error(fmt.Sprintf("[CreateCustomer] Error creating customer, %v", err))
		return nil, err
	}
	return newCustomer, nil
}

func (svc customerService) UpdateCustomer(customer models.Customer) (*models.Customer, error) {
	updatedCustomer, err := svc.model.UpdateCustomer(customer)
	if err != nil {
		logrus.Error(fmt.Sprintf("[UpdateCustomer] Error updating customer, %v", err))
		return nil, err
	}
	return updatedCustomer, nil
}

func (svc customerService) DeleteCustomer(ID int) error {
	var customerID uint
	customerID = uint(ID)
	err := svc.model.DeleteCustomer(customerID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[DeleteCustomer] Error deleting customer, %v", err))
		return err
	}
	return nil
}
