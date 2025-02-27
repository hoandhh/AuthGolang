package routes

import (
	"Auth/controllers"
	"Auth/middleware"

	"github.com/gin-gonic/gin"
)

var authController = new(controllers.AuthController) // Dùng khai báo được ở cấp package
var employeeController = new(controllers.EmployeeController)
var departmentController = new(controllers.DepartmentController)

func AuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}
}

func EmployeeRoutes(router *gin.RouterGroup) {
	employeeRoutes := router.Group("/employees").Use(middleware.AuthRequired())
	{
		employeeRoutes.POST("/", employeeController.CreateEmployee)
		employeeRoutes.PUT("/:id", employeeController.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeController.DeleteEmployee)
		employeeRoutes.POST("/assign-department", employeeController.AssignEmployee)
	}
}

func DepartmentRoutes(router *gin.RouterGroup) {
	departmentRoutes := router.Group("/departments").Use(middleware.AuthRequired())
	{
		departmentRoutes.POST("/", departmentController.CreateDepartment)
		departmentRoutes.PUT("/:id", departmentController.UpdateDepartment)
		departmentRoutes.DELETE("/:id", departmentController.DeleteDepartment)
	}
}
