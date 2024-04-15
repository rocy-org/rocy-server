package init

import (
	"github.com/rocy-org/rocy-server/database/album"
	"github.com/rocy-org/rocy-server/database/artist"
	"github.com/rocy-org/rocy-server/database/music"
	"github.com/rocy-org/rocy-server/database/playlist"
	"github.com/rocy-org/rocy-server/database/resource"
	"github.com/rocy-org/rocy-server/database/role"
	"github.com/rocy-org/rocy-server/database/user"
	"github.com/rocy-org/rocy-server/utils/db"
)

func init() {

	db.DB.AutoMigrate(
		&resource.Resource{},
		&artist.Artist{},
		&album.Album{},
		&music.Music{},
		&playlist.Playlist{},
		&role.Role{},
		&user.User{},
	)
}
