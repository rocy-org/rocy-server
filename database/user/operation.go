package user

import (
	"github.com/rocy-org/rocy-server/database/album"
	"github.com/rocy-org/rocy-server/utils/db"
)

var DB = db.DB

func (u *User) GetAllUser() ([]User, error) {
	var data []User
	err := DB.Model(&User{}).Preload("Album").Preload("Artist").Find(&data).Error
	return data, err
}

func (u *User) GetUserAlbum() ([]album.Album, error) {
	return nil, nil
}
