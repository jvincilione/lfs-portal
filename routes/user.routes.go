package routes

import (
	"lfs-portal/controllers"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func userRoutes(route *gin.RouterGroup, db *gorm.DB) {
	userModel := models.NewUser(db)
	userService := services.NewUserService(userModel)
	userController := controllers.NewUserController(userService)

	userRoutes := route.Group("/user")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/authenticate", userController.AuthenticateUser)
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserById)
		userRoutes.PATCH("/:id", userController.UpdateUser)
		userRoutes.PATCH("/:id/password", userController.UpdateUserPassword)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
