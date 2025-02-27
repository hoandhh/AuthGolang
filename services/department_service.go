package services

import (
	"Auth/database"
	"Auth/dtos"
	"Auth/models"
	"errors"
)

func CreateDepartment(dto dtos.DepartmentDTO) (models.Department, error) {
	department := models.Department{
		Name: dto.Name,
	}

	// INSERT INTO departments (name, created_at, updated_at) VALUES (?, ?, ?)
	if err := database.DB.Create(&department).Error; err != nil {
		return models.Department{}, err
	}

	return department, nil
}

func UpdateDepartment(id string, dto dtos.DepartmentDTO) (models.Department, error) {
	var department models.Department
	// SELECT * FROM departments WHERE id = ? LIMIT 1
	if err := database.DB.First(&department, id).Error; err != nil {
		return models.Department{}, errors.New("không tìm thấy phòng ban")
	}

	department.Name = dto.Name

	// UPDATE departments SET name = ?, updated_at = ? WHERE id = ?
	if err := database.DB.Save(&department).Error; err != nil {
		return models.Department{}, err
	}

	return department, nil
}

func DeleteDepartment(id string) error {
	// DELETE FROM departments WHERE id = ?
	if err := database.DB.Delete(&models.Department{}, id).Error; err != nil {
		return err
	}
	return nil
}
