package routes

import (
	"github.com/gin-gonic/gin"
)

func userRoutes(route *gin.RouterGroup) {

	userRoutes := route.Group("/user")
	{
		userRoutes.POST("/", UserController.CreateUser)
		userRoutes.POST("/authenticate", UserController.AuthenticateUser)
		userRoutes.POST("/logout", UserController.LogoutUser)
		userRoutes.GET("/", UserController.GetAllUsers)
		userRoutes.GET("/:id", UserController.GetUserById)
		userRoutes.GET("/:id/company", CompanyController.GetUserCompanies)
		userRoutes.PATCH("/:id", UserController.UpdateUser)
		userRoutes.PATCH("/:id/password", UserController.UpdateUserPassword)
		userRoutes.DELETE("/:id", UserController.DeleteUser)
	}
}
