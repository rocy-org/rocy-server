package artist

import "github.com/rocy-org/rocy-server/utils/db"

var DB = db.DB

func (a *Artist) GetAllArtist() ([]Artist, error) {
	var data []Artist
	err := DB.Model(&Artist{}).Find(&data).Error
	return data, err

}

func (a *Artist) GetArtistByName() error {
	return DB.Take(a, "name = ?", a.Name).Error
}
