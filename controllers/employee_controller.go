package controllers

import (
	"Auth/dtos"
	"Auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

// Thêm nhân sự
func (ctrl EmployeeController) CreateEmployee(c *gin.Context) {
	var employeeDTO dtos.EmployeeDTO
	if err := c.ShouldBindJSON(&employeeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	employee, err := services.CreateEmployee(employeeDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// Sửa nhân sự
func (ctrl EmployeeController) UpdateEmployee(c *gin.Context) {
	var employeeDTO dtos.EmployeeDTO
	if err := c.ShouldBindJSON(&employeeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	id := c.Param("id")
	updatedEmployee, err := services.UpdateEmployee(id, employeeDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

// Xóa nhân sự
func (ctrl EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "xoá nhân sự thành công"})
}

// Gán nhân sự vào phòng ban
func (ctrl EmployeeController) AssignEmployee(c *gin.Context) {
	var employeeDepartmentDTO dtos.EmployeeDepartmentDTO
	if err := c.ShouldBindJSON(&employeeDepartmentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := services.AssignEmployeeToDepartment(employeeDepartmentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "gán nhân sự vào phòng ban thành công"})
}
