package routes

import (
	"github.com/gin-gonic/gin"
)

func companyRoutes(route *gin.RouterGroup) {
	custRoutes := route.Group("/company")
	{
		custRoutes.POST("/", CompanyController.CreateCompany)
		custRoutes.GET("/", CompanyController.GetAllCompanies)
		custRoutes.GET("/:id", CompanyController.GetCompanyById)
		custRoutes.GET("/:id/job", JobController.GetCompanyJobs)
		custRoutes.GET("/:id/user", UserController.GetCompanyUsers)
		custRoutes.PATCH("/:id", CompanyController.UpdateCompany)
		custRoutes.DELETE("/:id", CompanyController.DeleteCompany)
	}
}
