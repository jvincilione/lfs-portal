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
	CustomerController interface {
		GetCustomerById(c *gin.Context)
		GetAllCustomers(c *gin.Context)
		CreateCustomer(c *gin.Context)
		UpdateCustomer(c *gin.Context)
		DeleteCustomer(c *gin.Context)
	}

	customerController struct {
		svc services.CustomerService
	}
)

func NewCustomerController(svc services.CustomerService) CustomerController {
	return customerController{svc}
}

func (controller customerController) GetCustomerById(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	customer, err := controller.svc.GetCustomerById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (controller customerController) GetAllCustomers(c *gin.Context) {
	customers, err := controller.svc.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": customers})
}

func (controller customerController) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(customer)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	newCustomer, err := controller.svc.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusCreated, newCustomer)
}

func (controller customerController) UpdateCustomer(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	customer, err := controller.svc.GetCustomerById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if customer.ID == 0 {
		logrus.Error(customer.ID)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	err = c.BindJSON(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(customer)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	updatedCustomer, err := controller.svc.UpdateCustomer(*customer)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusOK, updatedCustomer)
}

func (controller customerController) DeleteCustomer(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	err = controller.svc.DeleteCustomer(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
