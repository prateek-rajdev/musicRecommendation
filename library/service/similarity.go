package service

import "github.com/prateek-rajdev/musicRecommendation/library/model"

func NewSongSimilarityMap() *model.SongSimilarityMap {
	return &model.SongSimilarityMap{
		Similarity: make(map[string][]model.SongRecommendation),
	}
}
