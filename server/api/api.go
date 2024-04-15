package api

import (
	_ "github.com/rocy-org/rocy-server/docs/swagger"
)

// @title           rocy-server web API
// @version         1.0
// @description     This is the web API provided by rocy-server.
// @termsOfService  https://github.com/rocy-org/rocy-server

// @contact.name   API Support
// @contact.url    https://github.com/rocy-org/rocy-server

// @license.name  MIT
// @license.url   https://github.com/rocy-org/rocy-server/blob/master/LICENSE

// @host      localhost:35865
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// Routes config
const (
	// V1BasePath v1 api paths.
	V1BasePath = "/api/v1"

	// SwaggerDoc is the swagger document api path.
	SwaggerDoc = "/swagger/*any"

	// MusicBasePath defines music api paths.
	MusicBasePath     = "/music"
	MusicListAllMusic = ""

	// ArtistBasePath defines artist api paths.
	ArtistBasePath       = "/artist"
	ArtistListAllArtists = ""

	// AlbumBasePath defines album api paths.
	AlbumBasePath     = "/album"
	AlbumListAllAlbum = ""

	// PlaylistBasePath defines playlist api paths.
	PlaylistBasePath    = "/playlist"
	PlaylistAllPlaylist = ""
)
