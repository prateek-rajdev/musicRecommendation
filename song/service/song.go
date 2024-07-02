package service

import (
	"errors"
	"github.com/prateek-rajdev/musicRecommendation/dataStore"
	"github.com/prateek-rajdev/musicRecommendation/song/model"
	"github.com/prateek-rajdev/musicRecommendation/utils"
)

func AddSong(song model.Song) error {
	if _, exists := dataStore.GetLibrary().GetSong(song.Name); exists {
		return errors.New("song already exists")
	}
	dataStore.GetLibrary().AddSong(song)

	for _, existingSong := range dataStore.GetLibrary().GetAllSongs() {
		if existingSong.Name != song.Name {
			similarity := utils.CalculateSimilarityIndex(existingSong, song)
			dataStore.GetSongSimilarity().AddSimilarity(existingSong, song, similarity)
		}
	}
	return nil
}
