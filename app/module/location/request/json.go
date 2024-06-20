package request

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type LocationRequest struct {
	LocationID   uint64 `json:"location_id"`
	LocationName string `json:"location_name"`
	CreatedBy    string
	UpdatedBy    string
}

type LocationsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *LocationRequest) ToDomain() *schema.Location {
	return &schema.Location{
		LocationID:   req.LocationID,
		LocationName: req.LocationName,
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			UpdatedBy: req.UpdatedBy,
		},
	}
}
