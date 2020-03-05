package controllers

import (
	"fmt"
	"lfs-portal/models"
	"lfs-portal/services"
	"lfs-portal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type (
	CompanyController interface {
		GetCompanyById(c *gin.Context)
		GetAllCompanies(c *gin.Context)
		CreateCompany(c *gin.Context)
		UpdateCompany(c *gin.Context)
		DeleteCompany(c *gin.Context)
	}

	companyController struct {
		svc services.CompanyService
	}
)

func NewCompanyController(svc services.CompanyService) CompanyController {
	return companyController{svc}
}

func (controller companyController) GetCompanyById(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	company, err := controller.svc.GetCompanyById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (controller companyController) GetAllCompanies(c *gin.Context) {
	companies, err := controller.svc.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": companies})
}

func (controller companyController) CreateCompany(c *gin.Context) {
	var company models.Company
	err := c.BindJSON(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(company)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	newCompany, err := controller.svc.CreateCompany(company)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusCreated, newCompany)
}

func (controller companyController) UpdateCompany(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	company, err := controller.svc.GetCompanyById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if company.ID == 0 {
		logrus.Error(company.ID)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	err = c.BindJSON(company)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(company)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	updatedCompany, err := controller.svc.UpdateCompany(*company)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusOK, updatedCompany)
}

func (controller companyController) DeleteCompany(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	err = controller.svc.DeleteCompany(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
