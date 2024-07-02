package model

type Song struct {
	Name            string `json:"name" binding:"required"`
	Genre           string `json:"genre" binding:"required"`
	Tempo           string `json:"tempo" binding:"required"`
	Singer          string `json:"singer" binding:"required"`
	PopularityScore int    `json:"popularityScore" binding:"required"`
	ReleaseYear     int    `json:"releaseYear" binding:"required"`
}
