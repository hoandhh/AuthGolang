package main

import (
	"Auth/database"
	"Auth/middleware"
	"Auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	database.Connect()

	// Thiết lập routes
	v1 := r.Group("/api/v1")
	{
		routes.AuthRoutes(v1)
		routes.EmployeeRoutes(v1)
		routes.DepartmentRoutes(v1)
	}

	r.Run("0.0.0.0:8080") // Chạy server trên port 8080
}
