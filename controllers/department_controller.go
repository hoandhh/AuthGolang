package controllers

import (
	"Auth/dtos"
	"Auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct{}

// Thêm phòng ban
func (ctrl DepartmentController) CreateDepartment(c *gin.Context) {
	var departmentDTO dtos.DepartmentDTO
	if err := c.ShouldBindJSON(&departmentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	department, err := services.CreateDepartment(departmentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)
}

// Sửa phòng ban
func (ctrl DepartmentController) UpdateDepartment(c *gin.Context) {
	var departmentDTO dtos.DepartmentDTO
	if err := c.ShouldBindJSON(&departmentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	id := c.Param("id")
	updatedDepartment, err := services.UpdateDepartment(id, departmentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDepartment)
}

// Xóa phòng ban
func (ctrl DepartmentController) DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteDepartment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "xoá phòng ban thành công"})
}
