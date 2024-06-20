package controller

import "github.com/bangadam/go-fiber-starter/app/module/attendance/service"

type Controller struct {
	Attendance AttendanceController
}

func NewController(attendanceService service.AttendanceService) *Controller {
	return &Controller{
		Attendance: NewAttendanceController(attendanceService),
	}
}
