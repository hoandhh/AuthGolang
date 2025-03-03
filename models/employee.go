package models

type Employee struct {
	ID                  uint                 `gorm:"primaryKey; column:id; type:bigint(20) unsigned"`
	Email               string               `gorm:"unique; not null;column:email; type:varchar(255)"`
	Name                string               `gorm:"not null; column:name; type:varchar(255)"`
	Age                 int                  `gorm:"not null; column:age; type:int(11)"`
	Address             string               `gorm:"not null; column:address; type:text"`
	EmployeeDepartments []EmployeeDepartment `gorm:"foreignKey:EmployeeID"`
	Departments         []Department         `gorm:"many2many:employee_departments; foreignKey:ID; joinForeignKey:EmployeeID; References:ID; joinReferences:DepartmentID"`
}

func (Employee) TableName() string {
	return "employees"
}
