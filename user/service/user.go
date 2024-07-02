package service

import (
	"errors"
	"github.com/prateek-rajdev/musicRecommendation/dataStore"
	"github.com/prateek-rajdev/musicRecommendation/user/model"
)

func AddUser(user model.User) error {
	for songName := range user.Playlist {
		if _, exists := dataStore.GetLibrary().GetSong(songName); !exists {
			return errors.New("song " + songName + " not found in the library")
		}
	}

	for friendName := range user.Friends {
		if _, exists := dataStore.GetLibrary().GetUser(friendName); !exists {
			return errors.New("friend " + friendName + " not found in the library")
		}
	}

	if _, exists := dataStore.GetLibrary().GetUser(user.Name); exists {
		return errors.New("user " + user.Name + " already exists")
	}

	user.Friends = make(map[string]bool)
	dataStore.GetLibrary().AddUser(user)
	return nil
}

func AddFriend(userName, friendName string) error {
	user, userExists := dataStore.GetLibrary().GetUser(userName)
	_, friendExists := dataStore.GetLibrary().GetUser(friendName)

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
	dataStore.GetLibrary().AddUser(user)
	return nil
}
