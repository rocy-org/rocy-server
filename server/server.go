package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rocy-org/rocy-server/server/api"
	"github.com/rocy-org/rocy-server/server/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	serverDefaultAddr = ":35865"
)

// InitAndRun init and run the backend web server.
func InitAndRun() {
	// Controllers.
	c := controller.NewController()

	// Init gin instance.
	r := gin.Default()

	// Register routes.
	v1 := r.Group(api.V1BasePath)
	{
		music := v1.Group(api.MusicBasePath)
		{
			music.GET(api.MusicListAllMusic, c.ListAllMusic)
		}

		album := v1.Group(api.AlbumBasePath)
		{
			album.GET(api.AlbumListAllAlbum, c.ListAllAlbums)
		}

		artist := v1.Group(api.ArtistBasePath)
		{
			artist.GET(api.ArtistListAllArtists, c.ListAllArtists)
		}

		playlist := v1.Group(api.PlaylistBasePath)
		{
			playlist.GET(api.PlaylistAllPlaylist, c.ListAllPlaylists)
		}
	}

	// Register swag doc handler.
	r.GET(api.SwaggerDoc, ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(serverDefaultAddr); err != nil {
		fmt.Println("error running rocy-server web server: ", err)
	}
}
