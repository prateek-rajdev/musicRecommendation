package main

import (
	"github.com/prateek-rajdev/musicRecommendation/config"
	"github.com/prateek-rajdev/musicRecommendation/routers"
)

func main() {
	config.InitLibrary()
	r := routers.SetupRouter()
	r.Run(":8080")
}
