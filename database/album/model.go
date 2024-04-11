package album

import (
	"github.com/rocy-org/rocy-server/database/artist"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title    string        `gorm:"not null"`
	ArtistID int           `gorm:"not null"`
	Artist   artist.Artist `gorm:"foreignKey:ArtistID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Album) TableName() string {
	return "album"
}
