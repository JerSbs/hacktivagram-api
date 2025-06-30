package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/models"
	"p2-livecode-3-JerSbs/repository"
)

func CreatePostService(input dto.CreatePostRequest, userID uint) (*models.Post, error) {
	content := input.Content

	// Jika kosong → ambil konten dari API jokes
	if content == "" {
		joke, err := getRandomJoke()
		if err != nil {
			return nil, err
		}
		content = joke
	}

	post := &models.Post{
		Content:  content,
		ImageURL: input.ImageURL,
		UserID:   userID,
	}

	// Simpan ke database
	err := repository.CreatePost(post)
	if err != nil {
		return nil, err
	}

	// Simpan log aktivitas
	desc := fmt.Sprintf("user create new POST with ID %d", post.ID)
	_ = repository.CreateUserLog(userID, desc)

	return post, nil
}

// Helper function untuk ambil jokes dari API
func getRandomJoke() (string, error) {
	url := os.Getenv("JOKES_API_URL")
	apiKey := os.Getenv("JOKES_API_KEY")

	if url == "" || apiKey == "" {
		return "", errors.New("jokes API URL or key not set in .env")
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Api-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("failed to fetch joke")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var jokeResp []map[string]string
	json.Unmarshal(body, &jokeResp)

	if len(jokeResp) == 0 {
		return "", errors.New("no joke returned")
	}
	return jokeResp[0]["joke"], nil
}
