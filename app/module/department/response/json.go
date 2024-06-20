package response

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Department struct {
	DepartmentID   uint64 `json:"department_id"`
	DepartmentName string `json:"department_name"`
}

func FromDomain(department *schema.Department) (res *Department) {
	if department != nil {
		res = &Department{
			DepartmentID:   department.DepartmentID,
			DepartmentName: department.DepartmentName,
		}
	}

	return res
}
