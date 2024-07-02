package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek-rajdev/musicRecommendation/song/model"
	"github.com/prateek-rajdev/musicRecommendation/song/service"
	"net/http"
)

func AddSong(c *gin.Context) {
	var err error
	var song model.Song
	if err = c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.AddSong(song)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
