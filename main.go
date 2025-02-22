package main

import (
	"Auth/database"
	"Auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()

	// // Tạo bảng nếu chưa có
	// database.DB.AutoMigrate(&models.User{})

	// Thiết lập routes
	v1 := r.Group("/api/v1")
	{
		routes.AuthRoutes(v1)
		routes.EmployeeRoutes(v1)
		routes.DepartmentRoutes(v1)
	}

	r.Run(":8080") // Chạy server trên port 8080
}
