package playlist

import "github.com/rocy-org/rocy-server/utils/db"

var DB = db.DB

func (p *Playlist) GetAllPlaylist() ([]Playlist, error) {
	var data []Playlist
	err := DB.Preload("Music").Find(&data).Error
	return data, err
}

func (p *Playlist) GetPlaylistByMusic() ([]Playlist, error) {
	var data []Playlist
	err := DB.Preload("Music").Where("music_id = ?", p.Music).Find(&data).Error
	return data, err
}
