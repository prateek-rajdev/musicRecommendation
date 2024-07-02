package routers

import (
	"github.com/gin-gonic/gin"
	c3 "github.com/prateek-rajdev/musicRecommendation/recommend/controller"
	c1 "github.com/prateek-rajdev/musicRecommendation/song/controller"
	c2 "github.com/prateek-rajdev/musicRecommendation/user/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	songGroup := r.Group("/song")
	{
		songGroup.POST("/", c1.AddSong)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/", c2.AddUser)
		userGroup.POST("/:userName/friend", c2.AddFriend)
	}

	recommendGroup := r.Group("/recommend")
	{
		recommendGroup.GET("/:userName", c3.RecommendSongs)
	}

	return r
}
