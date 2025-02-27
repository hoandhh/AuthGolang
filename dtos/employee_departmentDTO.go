package dtos

import "time"

type EmployeeDepartmentDTO struct {
	ID           uint       `json:"id"`
	EmployeeID   uint       `json:"employee_id"`
	DepartmentID uint       `json:"department_id"`
	Position     string     `json:"position"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date,omitempty"`
}
