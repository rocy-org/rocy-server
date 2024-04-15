package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/rocy-org/rocy-server/server/model/artist"
	"net/http"
)

// ListAllArtists godoc
//
//  @Summary		Get all artists info
//  @Description	fetch artists from database
//  @Tags			artist
//  @Accept			json
//  @Produce		json
//  @Success		200		{object}		[]model.Artist
//  @Router			/artist [get]
func (c *Controller) ListAllArtists(ctx *gin.Context) {
	artist1 := model.NewArtist("test-artist1", []string{"test-album1-1", "test-album1-2"})
	artist2 := model.NewArtist("test-artist2", []string{"test-album2-1", "test-album2-2"})
	ctx.JSON(http.StatusOK, []model.Artist{*artist1, *artist2})
}
