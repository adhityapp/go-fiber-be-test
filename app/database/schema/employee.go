package schema

import "github.com/bangadam/go-fiber-starter/utils/helpers"

type Employee struct {
	EmployeeID   uint64 `gorm:"primary_key;column:employee_id"`
	EmployeeCode string `gorm:"column:employee_code;unique;not null"`
	EmployeeName string `gorm:"column:employee_name;not null"`
	DepartmentID uint64 `gorm:"column:department_id;not null"`
	PositionID   uint64 `gorm:"column:position_id;not null"`
	Superior     uint64 `gorm:"column:superior;not null"`
	Password     string `gorm:"column:password;not null"`
	Base
}

// compare password
func (u *Employee) ComparePassword(password string) bool {
	return helpers.ValidateHash(password, u.Password)
}
