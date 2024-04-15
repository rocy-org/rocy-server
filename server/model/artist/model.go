package artist

// Artist model describes all info about an artist used in server.
type Artist struct {
	Name   string   `json:"name"`
	Albums []string `json:"albums"`
}

// NewArtist constructs an Artist instance with given parameters:
//
// * name: Name of artist.
// * albums: All albums published by artist.
func NewArtist(name string, albums []string) *Artist {
	return &Artist{
		name,
		albums,
	}
}
