package service

import (
	"errors"
	"github.com/prateek-rajdev/musicRecommendation/config"
	"github.com/prateek-rajdev/musicRecommendation/song/model"
	"github.com/prateek-rajdev/musicRecommendation/utils"
)

func AddSong(song model.Song) error {
	if _, exists := config.GetLibrary().GetSong(song.Name); exists {
		return errors.New("song already exists")
	}
	config.GetLibrary().AddSong(song)

	for _, existingSong := range config.GetLibrary().GetAllSongs() {
		if existingSong.Name != song.Name {
			similarity := utils.CalculateSimilarityIndex(existingSong, song)
			config.GetSongSimilarity().AddSimilarity(existingSong, song, similarity)
		}
	}
	return nil
}
