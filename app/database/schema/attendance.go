package schema

import "time"

type Attendance struct {
	AttendanceID uint64    `gorm:"primary_key;column:attendance_id"`
	EmployeeID   uint64    `gorm:"column:employee_id"`
	LocationID   uint64    `gorm:"column:location_id"`
	AbsenIN      time.Time `gorm:"column:absen_in"`
	AbsenOUT     time.Time `gorm:"column:absen_out"`
	Base
}

type Report struct {
	AttendanceID   uint64
	Date           time.Time
	EmployeeCode   string //employee
	EmployeeName   string //employee
	DepartmentName string //employee join departemen
	PositionName   string //employee join position
	LocationName   string // join location
	AbsenIN        time.Time
	AbsenOUT       time.Time
}
