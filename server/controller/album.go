package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/rocy-org/rocy-server/server/model/album"
	"net/http"
)

// ListAllAlbums godoc
//
//  @Summary		Get all album info
//  @Description	fetch albums from database
//  @Tags			album
//  @Accept			json
//  @Produce		json
//  @Success		200		{object}	[]model.Album
//  @Router			/album [get]
func (c *Controller) ListAllAlbums(ctx *gin.Context) {
	album1 := model.NewAlbum("test-album1", []string{"test-artist1-1", "test-artist1-2"}, []string{"test-music1-1", "test-music1-2"})
	album2 := model.NewAlbum("test-album2", []string{"test-artist2-1", "test-artist2-2"}, []string{"test-music2-1", "test-music2-2"})
	ctx.JSON(http.StatusOK, []model.Album{*album1, *album2})

}
