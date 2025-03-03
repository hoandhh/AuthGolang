package models

import (
	"time"
)

type EmployeeDepartment struct {
	ID           uint       `gorm:"primaryKey; column:id; type:bigint(20) unsigned"`
	EmployeeID   uint       `gorm:"not null; column:employee_id; type:bigint(20) unsigned"`
	DepartmentID uint       `gorm:"not null; column:department_id; type:bigint(20) unsigned"`
	Position     string     `gorm:"not null; column:position; type:varchar(255)"`
	StartDate    time.Time  `gorm:"not null; column:start_date; type:date"`
	EndDate      *time.Time `gorm:"default:null;column:end_date; type:date"`
	Employee     Employee   `gorm:"foreignKey:EmployeeID"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
}

func (EmployeeDepartment) TableName() string {
	return "employee_departments"
}
