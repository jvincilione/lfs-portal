package controllers

import (
	"lfs-portal/models"
	"lfs-portal/services"
	"lfs-portal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type (
	UserController interface {
		GetUserById(c *gin.Context)
		GetAllUsers(c *gin.Context)
		GetCompanyUsers(c *gin.Context)
		CreateUser(c *gin.Context)
		UpdateUser(c *gin.Context)
		DeleteUser(c *gin.Context)
		AuthenticateUser(c *gin.Context)
		LogoutUser(c *gin.Context)
		UpdateUserPassword(c *gin.Context)
	}

	userController struct {
		svc services.UserService
	}
)

func NewUserController(svc services.UserService) UserController {
	return userController{svc}
}

func (controller userController) GetUserById(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	user, err := controller.svc.GetUserById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, utils.GetNotFoundError())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller userController) GetAllUsers(c *gin.Context) {
	users, err := controller.svc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": users})
}

func (controller userController) GetCompanyUsers(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	users, err := controller.svc.GetCompanyUsers(uint(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": users})
}

func (controller userController) CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	newUser, err := controller.svc.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (controller userController) UpdateUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	user, err := controller.svc.GetUserById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if user.ID == 0 {
		logrus.Error(user.ID)
		c.JSON(http.StatusNotFound, utils.GetNotFoundError())
		return
	}

	err = c.BindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	updatedUser, err := controller.svc.UpdateUser(*user)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (controller userController) DeleteUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	err = controller.svc.DeleteUser(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (controller userController) AuthenticateUser(c *gin.Context) {
	user := models.AuthUser{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	token := controller.svc.AuthenticateUser(user.Email, user.Password)
	if token == "" {
		c.JSON(http.StatusUnauthorized, utils.GetUnauthorizedError())
		return
	}

	expirationTime := int(time.Now().Add(120 * time.Minute).Unix())

	c.SetCookie("lfs-auth-token", token, expirationTime, "/", utils.Config.DOMAIN, http.SameSiteNoneMode, false, false)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller userController) LogoutUser(c *gin.Context) {
	services.LogoutUser(c)
	c.JSON(http.StatusOK, gin.H{})
}

func (controller userController) UpdateUserPassword(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	usr, err := controller.svc.GetUserById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}
	if usr.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var user models.AuthUser
	c.BindJSON(&user)
	err = controller.svc.UpdatePasswordForUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
