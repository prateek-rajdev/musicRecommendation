package model

type User struct {
	Name     string          `json:"name" binding:"required"`
	Playlist map[string]bool `json:"playlist,omitempty" binding:"required"`
	Friends  map[string]bool `json:"friends,omitempty"`
}
