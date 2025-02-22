package dtos

type DepartmentDTO struct {
	Name string `json:"name" binding:"required"`
}
