package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseURL := "http://localhost:8080"

	// Add Songs
	songs := []map[string]interface{}{
		{"name": "Song1", "genre": "Pop", "tempo": "Slow", "singer": "Singer1", "popularityScore": 8, "releaseYear": 2022},
		{"name": "Song2", "genre": "Rock", "tempo": "Fast", "singer": "Singer2", "popularityScore": 7, "releaseYear": 2021},
		{"name": "Song3", "genre": "Jazz", "tempo": "Slow", "singer": "Singer3", "popularityScore": 8, "releaseYear": 2020},
		{"name": "Song4", "genre": "Classical", "tempo": "Slow", "singer": "Singer4", "popularityScore": 9, "releaseYear": 2019},
		{"name": "Song5", "genre": "Pop", "tempo": "Fast", "singer": "Singer5", "popularityScore": 8, "releaseYear": 2023},
		// below cases should fail
		{"name": "Song1", "genre": "Pop", "tempo": "Slow", "singer": "Singer1", "popularityScore": 8.0, "releaseYear": 2022},
		{"name": "Song21", "genre": "Rock", "tempo": "Fast", "singer": "Singer2", "popularityScore": 7},
		{"name": "Song55", "tempo": "Fast", "singer": "Singer5", "popularityScore": 8, "releaseYear": 2023},
	}

	for _, song := range songs {
		addSong(baseURL, song)
	}

	// Add Users without friends
	users := []map[string]interface{}{
		{"name": "User1", "playlist": map[string]bool{"Song1": true}},
		{"name": "User2", "playlist": map[string]bool{"Song2": true, "Song3": true}},
		// below cases should fail
		{"name": "User1", "playlist": map[string]bool{"Song1111": true}},
	}

	// Add friends
	friends := []map[string]string{
		{"userName": "User1", "friend": "User2"},
		{"userName": "User2", "friend": "User1"},
		// below case should faild
		{"userName": "User7", "friend": "User1"},
	}

	for _, user := range users {
		addUser(baseURL, user)
	}

	for _, friend := range friends {
		addFriend(baseURL, friend["userName"], friend["friend"])
	}

	// Get Recommendations
	getRecommendations(baseURL, "User1")
	getRecommendations(baseURL, "User2")
}

func addSong(baseURL string, song map[string]interface{}) {
	jsonData, _ := json.Marshal(song)
	resp, err := http.Post(baseURL+"/song", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error adding song:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Add Song Response:", string(body))
}

func addUser(baseURL string, user map[string]interface{}) {
	jsonData, _ := json.Marshal(user)
	resp, err := http.Post(baseURL+"/user", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error adding user:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Add User Response:", string(body))
}

func addFriend(baseURL, userName, friendName string) {
	data := map[string]string{"friend": friendName}
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", baseURL+"/user/"+userName+"/friend", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error adding friend:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Add Friend Response:", string(body))
}

func getRecommendations(baseURL, userName string) {
	resp, err := http.Get(baseURL + "/recommend/" + userName)
	if err != nil {
		fmt.Println("Error getting recommendations:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Recommendations for", userName+":", string(body))
}
