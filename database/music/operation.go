package music

import (
	"github.com/rocy-org/rocy-server/utils/db"
)

var DB = db.DB

func (m *Music) GetAllMusic() ([]Music, error) {
	var data []Music
	err := DB.Model(&Music{}).
		Preload("Album").
		Preload("Artist").
		Preload("Resource").
		Find(&data).Error
	return data, err
}

func (m *Music) GetMusicByTitle() error {
	return DB.Model(&Music{}).
		Preload("Album").
		Preload("Artist").
		Preload("Resource").
		Take(m, "title = ? ", m.Title).Error
}

func (m *Music) GetMusicByAlbum(album string) ([]Music, error) {
	var data []Music
	err := DB.Model(&Music{}).
		Preload("Album").
		Preload("Artist").
		Preload("Resource").
		Where("Album = ?", album).Find(&data).Error

	return data, err
}

func (m *Music) GetMusicByResource(resource string) ([]Music, error) {
	// 通过 resource 查询 id
	var data []Music
	err := DB.Model(&Music{}).
		Preload("Album").
		Preload("Artist").
		Preload("Resource").
		Where("resource_id = ?", resource).
		Find(&data).Error

	return data, err

}
