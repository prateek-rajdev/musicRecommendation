package service

import (
	"errors"
	"github.com/prateek-rajdev/musicRecommendation/dataStore"
	model2 "github.com/prateek-rajdev/musicRecommendation/library/model"
	"sort"
)

func RecommendSongs(userName string) ([]model2.SongRecommendation, error) {
	user, exists := dataStore.GetLibrary().GetUser(userName)
	if !exists {
		return nil, errors.New("user not found")
	}

	recommendations := make([]model2.SongRecommendation, 0)
	resultSet := make(map[string]bool)
	var results []model2.SongRecommendation

	for songName, _ := range user.Playlist {
		similarSongs, exist := dataStore.GetSongSimilarity().GetSimilarity(songName)
		if exist {
			for _, s := range similarSongs {
				if _, exists = user.Playlist[s.Song.Name]; !exists {
					recommendations = append(recommendations, s)
				}
			}
		}
	}

	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	// send unique results
	for _, recommendation := range recommendations {
		if _, exists = resultSet[recommendation.Song.Name]; exists {
			continue
		}
		results = append(results, recommendation)
		resultSet[recommendation.Song.Name] = true
	}

	return results, nil
}
