package dataStore

import (
	"github.com/prateek-rajdev/musicRecommendation/library/model"
	"github.com/prateek-rajdev/musicRecommendation/library/service"
)

var Library *model.Library
var SongSimilarityMap *model.SongSimilarityMap

func InitLibrary() {
	Library = service.NewLibrary()
	SongSimilarityMap = service.NewSongSimilarityMap()
}

func GetLibrary() *model.Library {
	return Library
}

func GetSongSimilarity() *model.SongSimilarityMap {
	return SongSimilarityMap
}
