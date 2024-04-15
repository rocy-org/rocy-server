package user

import (
	"github.com/rocy-org/rocy-server/database/album"
	"github.com/rocy-org/rocy-server/database/artist"
	"github.com/rocy-org/rocy-server/database/role"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string          `gorm:"size:255;not null"`
	Password string          `gorm:"size:255;not null"`
	RoleID   int             `gorm:"not null"`
	Role     role.Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Album    []album.Album   `gorm:"many2many:user_album"`
	Artist   []artist.Artist `gorm:"many2many:user_artist"`
}

func (User) TableName() string {
	return "user"
}
