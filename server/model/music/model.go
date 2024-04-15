package music

import "time"

// Music model describes all info about a music instance used in server.
type Music struct {
	Title    string        `json:"title"`
	Artist   []string      `json:"artist"`
	Album    string        `json:"album"`
	Duration time.Duration `json:"duration"`
	Url      string        `json:"url"`
}

// NewMusic constructs a Music instance with given parameters:
//
// * title: Music title.
// * artist: Artist names, a list of string.
// * album: Album title.
// * duration: Music duration in milliseconds.
// * url: Music source url.
func NewMusic(title string, artist []string, album string, duration time.Duration, url string) *Music {
	return &Music{
		title,
		artist,
		album,
		duration,
		url,
	}
}
