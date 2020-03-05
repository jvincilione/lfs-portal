package routes

import (
	"lfs-portal/controllers"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func companyRoutes(route *gin.RouterGroup, db *gorm.DB) {
	companyModel := models.NewCompany(db)
	companyService := services.NewCompanyService(companyModel)
	companyController := controllers.NewCompanyController(companyService)

	custRoutes := route.Group("/company")
	{
		custRoutes.POST("/", companyController.CreateCompany)
		custRoutes.GET("/", companyController.GetAllCompanies)
		custRoutes.GET("/:id", companyController.GetCompanyById)
		custRoutes.PATCH("/:id", companyController.UpdateCompany)
		custRoutes.DELETE("/:id", companyController.DeleteCompany)
	}
}
