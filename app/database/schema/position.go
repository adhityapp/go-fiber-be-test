package schema

type Position struct {
	PositionID   uint64 `gorm:"primary_key;column:position_id"`
	DepartmentID uint64 `gorm:"column:department_id;not null"`
	PositionName string `gorm:"column:position_name;not null"`
	Base
}
