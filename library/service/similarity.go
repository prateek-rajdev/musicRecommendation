package service

import "musicRecommendation/library/model"

func NewSongSimilarityMap() *model.SongSimilarityMap {
	return &model.SongSimilarityMap{
		Similarity: make(map[string][]model.SongRecommendation),
	}
}
