package playlist

import (
	"github.com/rocy-org/rocy-server/database/music"
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Title   string `gorm:"not null"`
	MusicID int
	Music   music.Music `gorm:"foreignKey:MusicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Playlist) TableName() string {
	return "playlist"
}
