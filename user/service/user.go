package service

import (
	"errors"
	"github.com/prateek-rajdev/musicRecommendation/config"
	"github.com/prateek-rajdev/musicRecommendation/user/model"
)

func AddUser(user model.User) error {
	for songName := range user.Playlist {
		if _, exists := config.GetLibrary().GetSong(songName); !exists {
			return errors.New("song " + songName + " not found in the library")
		}
	}

	for friendName := range user.Friends {
		if _, exists := config.GetLibrary().GetUser(friendName); !exists {
			return errors.New("friend " + friendName + " not found in the library")
		}
	}

	if _, exists := config.GetLibrary().GetUser(user.Name); exists {
		return errors.New("user " + user.Name + " already exists")
	}

	user.Friends = make(map[string]bool)
	config.GetLibrary().AddUser(user)
	return nil
}

func AddFriend(userName, friendName string) error {
	user, userExists := config.GetLibrary().GetUser(userName)
	_, friendExists := config.GetLibrary().GetUser(friendName)

	if !userExists {
		return errors.New("user not found")
	}

	if !friendExists {
		return errors.New("friend not found")
	}

	if areFriends, exists := user.Friends[friendName]; exists && areFriends {
		return errors.New("users are already friends")
	}

	user.Friends[friendName] = true
	config.GetLibrary().AddUser(user)
	return nil
}
