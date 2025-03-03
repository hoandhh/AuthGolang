package models

type Department struct {
	ID                  uint                 `gorm:"primaryKey; column:id; type:bigint(20) unsigned"`
	Name                string               `gorm:"unique; not null; column:name; type:varchar(255)"`
	EmployeeDepartments []EmployeeDepartment `gorm:"foreignKey:DepartmentID"`
	Employees           []Employee           `gorm:"many2many:employee_departments; foreignKey:ID; joinForeignKey:DepartmentID; References:ID; joinReferences:EmployeeID"`
}

// TableName xác định tên bảng trong CSDL
func (Department) TableName() string {
	return "departments"
}
