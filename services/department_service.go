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

	if err := database.DB.Create(&department).Error; err != nil {
		return models.Department{}, err
	}

	return department, nil
}

func UpdateDepartment(id string, dto dtos.DepartmentDTO) (models.Department, error) {
	var department models.Department
	if err := database.DB.First(&department, id).Error; err != nil {
		return models.Department{}, errors.New("không tìm thấy phòng ban")
	}

	department.Name = dto.Name

	if err := database.DB.Save(&department).Error; err != nil {
		return models.Department{}, err
	}

	return department, nil
}

func DeleteDepartment(id string) error {
	if err := database.DB.Delete(&models.Department{}, id).Error; err != nil {
		return err
	}
	return nil
}
