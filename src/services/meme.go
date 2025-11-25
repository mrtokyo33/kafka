package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kafka/src/models"
)

func GetRandomMeme(subreddit string) (*models.MemeResponse, error) {
	url := "https://meme-api.com/gimme"
	if subreddit != "" {
		url = fmt.Sprintf("https://meme-api.com/gimme/%s", subreddit)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var meme models.MemeResponse
	if err := json.NewDecoder(resp.Body).Decode(&meme); err != nil {
		return nil, err
	}

	return &meme, nil
}
