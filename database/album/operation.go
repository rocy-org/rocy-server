package album

import "github.com/rocy-org/rocy-server/utils/db"

var DB = db.DB

func (a *Album) GetAlbumByTitle() error {
	return DB.Model(&Album{}).
		Preload("Artist").
		Take("title = ?", a.Title).Error
}

func (a *Album) GetAlbumByArtist() ([]Album, error) {
	var data []Album
	err := DB.Model(&Album{}).
		Preload("Artist").
		Where("artist_id = ?", a.Artist).
		Find(&data).Error
	return data, err
}

func (a *Album) GetAllAlbum() ([]Album, error) {
	var data []Album
	err := DB.Model(&Album{}).
		Preload("Artist").
		Find(&data).Error
	return data, err

}