package music

import (
	"github.com/rocy-org/rocy-server/database/album"
	"github.com/rocy-org/rocy-server/database/artist"
	"github.com/rocy-org/rocy-server/database/resource"
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Title      string            `gorm:"not null"`
	Url        string            `gorm:"not null"`
	AlbumID    int               `gorm:"not null"`
	Album      album.Album       `gorm:"foreignKey:AlbumID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArtistID   int               `gorm:"not null"`
	Artist     artist.Artist     `gorm:"foreignKey:ArtistID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ResourceID int               `gorm:"not null"`
	Resource   resource.Resource `gorm:"foreignKey:ResourceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Music) TableName() string {
	return "music"
}
