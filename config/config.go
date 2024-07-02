package config

import (
	"musicRecommendation/library/model"
	"musicRecommendation/library/service"
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
