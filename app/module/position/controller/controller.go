package controller

import "github.com/bangadam/go-fiber-starter/app/module/position/service"

type Controller struct {
	Position PositionController
}

func NewController(positionService service.PositionService) *Controller {
	return &Controller{
		Position: NewPositionController(positionService),
	}
}
