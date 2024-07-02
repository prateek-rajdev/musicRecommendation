package service

import (
	model3 "github.com/prateek-rajdev/musicRecommendation/library/model"
	"github.com/prateek-rajdev/musicRecommendation/song/model"
	model2 "github.com/prateek-rajdev/musicRecommendation/user/model"
)

func NewLibrary() *model3.Library {
	return &model3.Library{
		Songs: make(map[string]model.Song),
		Users: make(map[string]model2.User),
	}
}
