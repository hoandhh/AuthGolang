package dtos

type EmployeeDTO struct {
	Email   string `json:"email" binding:"required,email"`
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
}
