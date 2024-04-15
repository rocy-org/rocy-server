package playlist

type Playlist struct {
	Name  string   `json:"name"`
	Music []string `json:"music"`
}

// NewPlaylist constructs a Playlist instance with given parameters:
//
// * name: Playlist name.
// * music: All music in playlist.
func NewPlaylist(name string, music []string) *Playlist {
	return &Playlist{
		name,
		music,
	}
}
