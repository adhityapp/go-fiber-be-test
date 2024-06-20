package controller

import "github.com/bangadam/go-fiber-starter/app/module/report/service"

type Controller struct {
	Report ReportController
}

func NewController(reportService service.ReportService) *Controller {
	return &Controller{
		Report: NewReportController(reportService),
	}
}
