package model

import (
	"github.com/prateek-rajdev/musicRecommendation/song/model"
	model2 "github.com/prateek-rajdev/musicRecommendation/user/model"
)

type Library struct {
	Songs map[string]model.Song
	Users map[string]model2.User
}

// TODO: make it scalable using matrices
func (lib *Library) AddSong(song model.Song) {
	lib.Songs[song.Name] = song
}

func (lib *Library) AddUser(user model2.User) {
	lib.Users[user.Name] = user
}

func (lib *Library) GetUser(userName string) (model2.User, bool) {
	if data, exist := lib.Users[userName]; exist {
		return data, true
	}
	return model2.User{}, false
}

func (lib *Library) GetSong(songName string) (model.Song, bool) {
	if data, exist := lib.Songs[songName]; exist {
		return data, true
	}
	return model.Song{}, false
}

func (lib *Library) GetAllSongs() map[string]model.Song {
	return lib.Songs
}
