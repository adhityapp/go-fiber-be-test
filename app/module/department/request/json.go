package request

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type DepartmentRequest struct {
	DepartmentID   uint64 `json:"department_id"`
	DepartmentName string `json:"department_name"`
	CreatedBy      string
	UpdatedBy      string
}

type DepartmentsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *DepartmentRequest) ToDomain() *schema.Department {
	return &schema.Department{
		DepartmentID:   req.DepartmentID,
		DepartmentName: req.DepartmentName,
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			UpdatedBy: req.UpdatedBy,
		},
	}
}
