package models

import "time"

type EmployeeDepartment struct {
	ID           uint       `gorm:"primaryKey"`
	EmployeeID   uint       `gorm:"not null"`
	DepartmentID uint       `gorm:"not null"`
	Position     string     `gorm:"not null"`
	StartDate    time.Time  `gorm:"not null"`
	EndDate      *time.Time `gorm:"default:null"`
}
