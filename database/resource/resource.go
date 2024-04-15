package resource

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Url    string `gorm:"not null"`
	Online bool   `gorm:"not null"`
}

func (Resource) TableName() string {
	return "resource"
}
