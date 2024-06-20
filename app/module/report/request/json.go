package request

import (
	"time"

	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type ReportRequest struct {
	AttendanceID uint64    `json:"attendance_id"`
	EmployeeID   uint64    `json:"employee_id"`
	LocationID   uint64    `json:"location_id"`
	AbsenIN      time.Time `json:"absen_in"`
	AbsenOUT     time.Time `json:"absen_out"`
	CreatedBy    string
	UpdatedBy    string
}

type ReportsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

// func (req *ReportRequest) ToDomain() *schema.Attendance {
// 	return &schema.Attendance{
// 		AttendanceID: req.AttendanceID,
// 		EmployeeID:   req.EmployeeID,
// 		LocationID:   req.LocationID,
// 		AbsenIN:      req.AbsenIN,
// 		AbsenOUT:     req.AbsenOUT,
// 		Base: schema.Base{
// 			CreatedBy: req.CreatedBy,
// 			UpdatedBy: req.UpdatedBy,
// 		},
// 	}
// }
