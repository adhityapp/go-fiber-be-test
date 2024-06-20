package controller

import "github.com/bangadam/go-fiber-starter/app/module/department/service"

type Controller struct {
	Department DepartmentController
}

func NewController(departmentService service.DepartmentService) *Controller {
	return &Controller{
		Department: NewDepartmentController(departmentService),
	}
}
