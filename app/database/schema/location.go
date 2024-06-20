package schema

type Location struct {
	LocationID   uint64 `gorm:"primary_key;column:location_id"`
	LocationName string `gorm:"column:location_name;not null"`
	Base
}
