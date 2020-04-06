package routes

import (
	"lfs-portal/controllers"
	database "lfs-portal/db"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine

	JobModel      models.JobModel
	JobService    services.JobService
	JobController controllers.JobController

	CompanyModel      models.CompanyModel
	CompanyService    services.CompanyService
	CompanyController controllers.CompanyController

	UserModel      models.UserModel
	UserService    services.UserService
	UserController controllers.UserController
)

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
