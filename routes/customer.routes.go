package routes

import (
	"lfs-portal/controllers"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func customerRoutes(route *gin.RouterGroup, db *gorm.DB) {
	customerModel := models.NewCustomer(db)
	customerService := services.NewCustomerService(customerModel)
	customerController := controllers.NewCustomerController(customerService)

	custRoutes := route.Group("/customer")
	{
		custRoutes.POST("/", customerController.CreateCustomer)
		custRoutes.GET("/", customerController.GetAllCustomers)
		custRoutes.GET("/:id", customerController.GetCustomerById)
		custRoutes.PATCH("/:id", customerController.UpdateCustomer)
		custRoutes.DELETE("/:id", customerController.DeleteCustomer)
	}
}
