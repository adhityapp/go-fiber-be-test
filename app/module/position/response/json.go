package response

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Position struct {
	PositionID   uint64 `json:"column:position_id"`
	DepartmentID uint64 `json:"department_id"`
	PositionName string `json:"position_name"`
}

func FromDomain(position *schema.Position) (res *Position) {
	if position != nil {
		res = &Position{
			PositionID:   position.PositionID,
			DepartmentID: position.DepartmentID,
			PositionName: position.PositionName,
		}
	}

	return res
}
