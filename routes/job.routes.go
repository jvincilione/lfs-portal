package routes

import (
	"lfs-portal/controllers"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func jobRoutes(route *gin.RouterGroup, db *gorm.DB) {
	jobModel := models.NewJob(db)
	jobService := services.NewJobService(jobModel)
	jobController := controllers.NewJobController(jobService)

	custRoutes := route.Group("/job")
	{
		custRoutes.POST("/", jobController.CreateJob)
		custRoutes.GET("/", jobController.GetAllJobs)
		custRoutes.GET("/:id", jobController.GetJobById)
		custRoutes.PATCH("/:id", jobController.UpdateJob)
		custRoutes.DELETE("/:id", jobController.DeleteJob)
	}
}
