package model

import "musicRecommendation/song/model"

// SongRecommendation embeds Song and adds a similarity score.
type SongRecommendation struct {
	Song  model.Song `json:"song"`
	Score float64    `json:"score"`
}

// SongSimilarityMap stores recommendations with similarity scores.
type SongSimilarityMap struct {
	Similarity map[string][]SongRecommendation
}

func (ssm *SongSimilarityMap) AddSimilarity(song1, song2 model.Song, similarity float64) {
	if similarity > 0.0 {
		ssm.Similarity[song1.Name] = insertSorted(ssm.Similarity[song1.Name], SongRecommendation{Song: song2, Score: similarity})
		ssm.Similarity[song2.Name] = insertSorted(ssm.Similarity[song2.Name], SongRecommendation{Song: song1, Score: similarity})
	}
}

func (ssm *SongSimilarityMap) GetSimilarity(songName string) ([]SongRecommendation, bool) {
	var s []SongRecommendation
	if similarSongs, exists := ssm.Similarity[songName]; exists {
		return similarSongs, true
	}
	return s, false
}

func insertSorted(recommendations []SongRecommendation, newRecommendation SongRecommendation) []SongRecommendation {
	index := len(recommendations)
	for i, recommendation := range recommendations {
		if newRecommendation.Score > recommendation.Score {
			index = i
			break
		}
	}
	recommendations = append(recommendations, SongRecommendation{}) // Extend the slice
	copy(recommendations[index+1:], recommendations[index:])        // Shift elements to the right
	recommendations[index] = newRecommendation                      // Insert new element
	return recommendations
}
