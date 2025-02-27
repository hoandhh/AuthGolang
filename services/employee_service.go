package services

import (
	"Auth/database"
	"Auth/dtos"
	"Auth/models"
	"errors"
	"time"
)

func CreateEmployee(dto dtos.EmployeeDTO) (models.Employee, error) {
	employee := models.Employee{
		Email:   dto.Email,
		Name:    dto.Name,
		Age:     dto.Age,
		Address: dto.Address,
	}

	// INSERT INTO employees (email, name, age, address, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)
	if err := database.DB.Create(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func UpdateEmployee(id string, dto dtos.EmployeeDTO) (models.Employee, error) {
	var employee models.Employee
	// SELECT * FROM employees WHERE id = ? LIMIT 1
	if err := database.DB.First(&employee, id).Error; err != nil {
		return models.Employee{}, errors.New("không tìm thấy nhân sự")
	}

	employee.Email = dto.Email
	employee.Name = dto.Name
	employee.Age = dto.Age
	employee.Address = dto.Address

	// UPDATE employees SET email = ?, name = ?, age = ?, address = ?, updated_at = ? WHERE id = ?
	if err := database.DB.Save(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func DeleteEmployee(id string) error {
	// DELETE FROM employees WHERE id = ?
	if err := database.DB.Delete(&models.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Gán nhân sự vào phòng ban
func AssignEmployeeToDepartment(dto dtos.EmployeeDepartmentDTO) error {
	var count int64

	// Kiểm tra nhân sự có tồn tại không
	database.DB.Model(&models.Employee{}).Where("id = ?", dto.EmployeeID).Count(&count)
	if count == 0 {
		return errors.New("nhân sự không tồn tại")
	}

	// Kiểm tra phòng ban có tồn tại không
	database.DB.Model(&models.Department{}).Where("id = ?", dto.DepartmentID).Count(&count)
	if count == 0 {
		return errors.New("phòng ban không tồn tại")
	}

	// Tạo đối tượng EmployeeDepartment từ DTO
	employeeDepartment := models.EmployeeDepartment{
		EmployeeID:   dto.EmployeeID,
		DepartmentID: dto.DepartmentID,
		StartDate:    dto.StartDate,
	}

	// Gán start_date nếu chưa có
	if employeeDepartment.StartDate.IsZero() {
		employeeDepartment.StartDate = time.Now()
	}

	// Thêm dữ liệu vào bảng employee_departments
	if err := database.DB.Create(&employeeDepartment).Error; err != nil {
		return errors.New("lỗi khi lưu vào DB")
	}

	return nil
}
