package schema

type Department struct {
	DepartmentID   uint64 `gorm:"primary_key;column:department_id"`
	DepartmentName string `gorm:"column:department_name;not null"`
	Base
}
