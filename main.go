package main

import (
	"github.com/prateek-rajdev/musicRecommendation/dataStore"
	"github.com/prateek-rajdev/musicRecommendation/routers"
)

func main() {
	dataStore.InitLibrary()
	r := routers.SetupRouter()
	r.Run(":8080")
}
