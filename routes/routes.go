package routes

import (
	"lfs-portal/controllers"
	database "lfs-portal/db"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var JobModel models.JobModel
var JobService services.JobService
var JobController controllers.JobController

var CompanyModel models.CompanyModel
var CompanyService services.CompanyService
var CompanyController controllers.CompanyController

var UserModel models.UserModel
var UserService services.UserService
var UserController controllers.UserController

func InitializeRoutes() *gin.Engine {
	db := database.NewDb()
	router = gin.Default()
	router.Use(services.AuthenticationMiddleware)

	JobModel = models.NewJob(db)
	JobService = services.NewJobService(JobModel)
	JobController = controllers.NewJobController(JobService)

	CompanyModel = models.NewCompany(db)
	CompanyService = services.NewCompanyService(CompanyModel)
	CompanyController = controllers.NewCompanyController(CompanyService)

	UserModel = models.NewUser(db)
	UserService = services.NewUserService(UserModel)
	UserController = controllers.NewUserController(UserService)

	api := router.Group("/api")
	{
		companyRoutes(api)
		jobRoutes(api)
		userRoutes(api)
	}
	router.NoRoute(func(c *gin.Context) {
		c.File("./lfs-frontend/build")
	})
	router.Use(static.Serve("/", static.LocalFile("./lfs-frontend/build", true)))
	return router
}
