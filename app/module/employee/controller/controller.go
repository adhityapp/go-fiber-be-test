package controller

import "github.com/bangadam/go-fiber-starter/app/module/employee/service"

type Controller struct {
	Employee EmployeeController
}

func NewController(employeeService service.EmployeeService) *Controller {
	return &Controller{
		Employee: NewEmployeeController(employeeService),
	}
}
