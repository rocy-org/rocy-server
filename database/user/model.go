package user

import (
	"github.com/rocy-org/rocy-server/database/album"
	"github.com/rocy-org/rocy-server/database/artist"
	"github.com/rocy-org/rocy-server/database/role"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string        `gorm:"size:255;not null"`
	Password string        `gorm:"size:255;not null"`
	RoleID   int           `gorm:"not null"`
	Role     role.Role     `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AlbumID  int           `gorm:"not null"`
	Album    album.Album   `gorm:"foreignKey:AlbumID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArtistID int           `gorm:"not null"`
	Artist   artist.Artist `gorm:"foreignKey:ArtistID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "user"
}
