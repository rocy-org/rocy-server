package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/rocy-org/rocy-server/server/model/playlist"
	"net/http"
)

// ListAllPlaylists godoc
//
//  @Summary		Get all playlists info
//  @Description	fetch playlists from database
//  @Tags			playlist
//  @Accept			json
//  @Produce		json
//  @Success		200		{object}	[]model.Playlist
//  @Router			/playlist [get]
func (c *Controller) ListAllPlaylists(ctx *gin.Context) {
	playlist1 := model.NewPlaylist("test-playlist1", []string{"test-music1-1", "test-music1-2"})
	playlist2 := model.NewPlaylist("test-playlist2", []string{"test-music2-1", "test-music2-2"})
	ctx.JSON(http.StatusOK, []model.Playlist{*playlist1, *playlist2})
}
