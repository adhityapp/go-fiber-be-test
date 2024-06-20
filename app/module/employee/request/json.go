package request

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/utils/helpers"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type EmployeeRequest struct {
	EmployeeCode string
	EmployeeName string `json:"employee_name" validate:"required"`
	DepartmentID uint64 `json:"department_id" validate:"required"`
	PositionID   uint64 `json:"position_id" validate:"required"`
	Superior     uint64 `json:"superior" validate:"required"`
	Password     string `json:"password" validate:"required"`
	CreatedBy    string
	UpdatedBy    string
}

type EmployeesRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *EmployeeRequest) ToDomain() *schema.Employee {
	return &schema.Employee{
		EmployeeCode: req.EmployeeCode,
		EmployeeName: req.EmployeeName,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		Superior:     req.Superior,
		Password:     helpers.Hash([]byte(req.Password)),
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			UpdatedBy: req.UpdatedBy,
		},
	}
}
