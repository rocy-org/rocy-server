package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/rocy-org/rocy-server/server/model/music"
	"net/http"
	"time"
)

// ListAllMusic godoc
//
//  @Summary		Get all music info
//  @Description	fetch music from database.
//  @Tags			music
//  @Accept			json
//  @Produce		json
//  @Success		200		{object}	[]model.Music
//  @Router			/music [get]
func (c *Controller) ListAllMusic(ctx *gin.Context) {
	music1 := model.NewMusic("test-title2", []string{"test-artist1-1", "test-artist1-2"}, "test-album1", time.Duration(100)*time.Millisecond, "test-url1")
	music2 := model.NewMusic("test-title2", []string{"test-artist2-1", "test-artist2-2"}, "test-album2", time.Duration(200)*time.Millisecond, "test-url2")
	ctx.JSON(http.StatusOK, []model.Music{*music1, *music2})
}
