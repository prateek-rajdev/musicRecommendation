package controller

import (
	"github.com/gin-gonic/gin"
	"musicRecommendation/recommend/service"
	"net/http"
)

func RecommendSongs(c *gin.Context) {
	userName := c.Param("userName")
	recommendations, err := service.RecommendSongs(userName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}
