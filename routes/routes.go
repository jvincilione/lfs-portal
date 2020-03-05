package routes

import (
	database "lfs-portal/db"
	"lfs-portal/services"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitializeRoutes() *gin.Engine {
	db := database.NewDb()
	router = gin.Default()
	router.Use(services.AuthenticationMiddleware)
	api := router.Group("/api")
	{
		companyRoutes(api, db)
		jobRoutes(api, db)
		userRoutes(api, db)
	}
	return router
}
