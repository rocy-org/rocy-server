package role

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string
}

func (Role) TableName() string {
	return "role"
}
