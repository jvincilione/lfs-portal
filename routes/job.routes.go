package routes

import (
	"github.com/gin-gonic/gin"
)

func jobRoutes(route *gin.RouterGroup) {

	custRoutes := route.Group("/job")
	{
		custRoutes.POST("/", JobController.CreateJob)
		custRoutes.GET("/", JobController.GetAllJobs)
		custRoutes.GET("/:id", JobController.GetJobById)
		custRoutes.PATCH("/:id", JobController.UpdateJob)
		custRoutes.DELETE("/:id", JobController.DeleteJob)
	}
}
