package album

// Album model describes all info about an album used in server.
type Album struct {
	Title  string   `json:"title"`
	Artist []string `json:"artist"`
	Music  []string `json:"music"`
}

// NewAlbum constructs an Album instance with given parameters:
//
// * title: Album title.
// * artist: All artists in album.
// * music: All music in album.
func NewAlbum(title string, artist []string, music []string) *Album {
	return &Album{
		title,
		artist,
		music,
	}
}
