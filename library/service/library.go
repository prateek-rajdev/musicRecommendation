package service

import (
	model3 "musicRecommendation/library/model"
	"musicRecommendation/song/model"
	model2 "musicRecommendation/user/model"
)

func NewLibrary() *model3.Library {
	return &model3.Library{
		Songs: make(map[string]model.Song),
		Users: make(map[string]model2.User),
	}
}
