package services

import (
	"fmt"
	"lfs-portal/models"

	"github.com/sirupsen/logrus"
)

type (
	CompanyService interface {
		GetCompanyById(ID int) (*models.Company, error)
		GetAllCompanies() ([]models.Company, error)
		GetUserCompanies(userID uint) ([]models.Company, error)
		CreateCompany(company models.Company) (*models.Company, error)
		UpdateCompany(company models.Company) (*models.Company, error)
		DeleteCompany(ID int) error
	}

	companyService struct {
		model models.CompanyModel
	}
)

func NewCompanyService(model models.CompanyModel) CompanyService {
	return companyService{model}
}

func (svc companyService) GetCompanyById(ID int) (*models.Company, error) {
	company, err := svc.model.GetCompanyById(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetCompanyById] Error getting company, %v", err))
		return nil, err
	}
	return company, nil
}

func (svc companyService) GetAllCompanies() ([]models.Company, error) {
	companies, err := svc.model.GetAllCompanies()
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetAllCompanies] Error getting companies, %v", err))
		return nil, err
	}
	return companies, nil
}

func (svc companyService) GetUserCompanies(userId uint) ([]models.Company, error) {
	companies, err := svc.model.GetUserCompanies(userId)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetAllCompanies] Error getting companies, %v", err))
		return nil, err
	}
	return companies, nil
}

func (svc companyService) CreateCompany(company models.Company) (*models.Company, error) {
	newCompany, err := svc.model.CreateCompany(company)
	if err != nil {
		logrus.Error(fmt.Sprintf("[CreateCompany] Error creating company, %v", err))
		return nil, err
	}
	return newCompany, nil
}

func (svc companyService) UpdateCompany(company models.Company) (*models.Company, error) {
	updatedCompany, err := svc.model.UpdateCompany(company)
	if err != nil {
		logrus.Error(fmt.Sprintf("[UpdateCompany] Error updating company, %v", err))
		return nil, err
	}
	return updatedCompany, nil
}

func (svc companyService) DeleteCompany(ID int) error {
	var companyID uint
	companyID = uint(ID)
	err := svc.model.DeleteCompany(companyID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[DeleteCompany] Error deleting company, %v", err))
		return err
	}
	return nil
}
