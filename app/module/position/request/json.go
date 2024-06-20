package request

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type PositionRequest struct {
	PositionID   uint64 `json:"column:position_id"`
	DepartmentID uint64 `json:"department_id"`
	PositionName string `json:"position_name"`
	CreatedBy    string
	UpdatedBy    string
}

type PositionsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *PositionRequest) ToDomain() *schema.Position {
	return &schema.Position{
		PositionID:   req.PositionID,
		DepartmentID: req.DepartmentID,
		PositionName: req.PositionName,
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			UpdatedBy: req.UpdatedBy,
		},
	}
}
