package services

import (
	"Auth/database"
	"Auth/dtos"
	"Auth/models"
	"errors"
)

func CreateEmployee(dto dtos.EmployeeDTO) (models.Employee, error) {
	employee := models.Employee{
		Email:   dto.Email,
		Name:    dto.Name,
		Age:     dto.Age,
		Address: dto.Address,
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func UpdateEmployee(id string, dto dtos.EmployeeDTO) (models.Employee, error) {
	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		return models.Employee{}, errors.New("không tìm thấy nhân sự")
	}

	employee.Email = dto.Email
	employee.Name = dto.Name
	employee.Age = dto.Age
	employee.Address = dto.Address

	if err := database.DB.Save(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func DeleteEmployee(id string) error {
	if err := database.DB.Delete(&models.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}
