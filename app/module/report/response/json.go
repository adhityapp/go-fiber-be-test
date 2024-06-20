package response

import (
	"time"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Report struct {
	AttendanceID   uint64    `json:"attendance_id"`
	Date           time.Time `json:"date"`
	EmployeeCode   string    `json:"employee_code"`   //employee
	EmployeeName   string    `json:"employee_name"`   //employee
	DepartmentName string    `json:"department_name"` //employee join departemen
	PositionName   string    `json:"position_name"`   //employee join position
	LocationName   string    `json:"location_name"`   // join location
	AbsenIN        time.Time `json:"absen_in"`
	AbsenOUT       time.Time `json:"absen_out"`
}

func FromDomain(attendance *schema.Report) (res *Report) {
	if attendance != nil {
		res = &Report{
			AttendanceID:   attendance.AttendanceID,
			Date:           attendance.Date,
			EmployeeCode:   attendance.EmployeeCode,
			EmployeeName:   attendance.EmployeeName,
			DepartmentName: attendance.DepartmentName,
			PositionName:   attendance.PositionName,
			LocationName:   attendance.LocationName,
			AbsenIN:        attendance.AbsenIN,
			AbsenOUT:       attendance.AbsenOUT,
		}
	}

	return res
}
