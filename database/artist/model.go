package artist

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (Artist) TableName() string {
	return "artist"
}
