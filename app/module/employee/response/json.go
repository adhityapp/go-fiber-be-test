package response

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Employee struct {
	EmployeeID   uint64 `json:"employee_id"`
	EmployeeCode string `json:"employee_code"`
	EmployeeName string `json:"employee_name"`
	DepartmentID uint64 `json:"department_id"`
	PositionID   uint64 `json:"position_id"`
	Superior     uint64 `json:"superior"`
	Password     string `json:"password"`
}

func FromDomain(employee *schema.Employee) (res *Employee) {
	if employee != nil {
		res = &Employee{
			EmployeeID:   employee.EmployeeID,
			EmployeeCode: employee.EmployeeCode,
			EmployeeName: employee.EmployeeName,
			DepartmentID: employee.DepartmentID,
			PositionID:   employee.PositionID,
			Superior:     employee.Superior,
			Password:     employee.Password,
		}
	}

	return res
}
