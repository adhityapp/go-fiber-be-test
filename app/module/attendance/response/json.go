package response

import (
	"time"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Attendance struct {
	AttendanceID uint64    `json:"attendance_id"`
	EmployeeID   uint64    `json:"employee_id"`
	LocationID   uint64    `json:"location_id"`
	AbsenIN      time.Time `json:"absen_in"`
	AbsenOUT     time.Time `json:"absen_out"`
}

func FromDomain(attendance *schema.Attendance) (res *Attendance) {
	if attendance != nil {
		res = &Attendance{
			AttendanceID: attendance.AttendanceID,
			EmployeeID:   attendance.EmployeeID,
			LocationID:   attendance.LocationID,
			AbsenIN:      attendance.AbsenIN,
			AbsenOUT:     attendance.AbsenOUT,
		}
	}

	return res
}
