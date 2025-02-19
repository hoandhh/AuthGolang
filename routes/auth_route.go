package routes

import (
	"Auth/controllers"

	"github.com/gin-gonic/gin"
)

var authController = new(controllers.AuthController) // Dùng khai báo được ở cấp package

func AuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}
}
