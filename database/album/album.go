package album

import (
	"github.com/rocy-org/rocy-server/database/artist"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title  string          `gorm:"not null"`
	Artist []artist.Artist `gorm:"many2many:album_artist;"`
}

func (Album) TableName() string {
	return "album"
}
