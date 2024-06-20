package controller

import "github.com/bangadam/go-fiber-starter/app/module/location/service"

type Controller struct {
	Location LocationController
}

func NewController(locationService service.LocationService) *Controller {
	return &Controller{
		Location: NewLocationController(locationService),
	}
}
