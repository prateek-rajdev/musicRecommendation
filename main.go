package main

import (
	"musicRecommendation/config"
	"musicRecommendation/routers"
)

func main() {
	config.InitLibrary()
	r := routers.SetupRouter()
	r.Run(":8080")
}
