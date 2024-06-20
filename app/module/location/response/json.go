package response

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Location struct {
	LocationID   uint64 `json:"location_id"`
	LocationName string `json:"location_name"`
}

func FromDomain(location *schema.Location) (res *Location) {
	if location != nil {
		res = &Location{
			LocationID:   location.LocationID,
			LocationName: location.LocationName,
		}
	}

	return res
}
