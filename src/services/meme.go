package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"kafka/src/config"
	"kafka/src/models"
)

func cleanSubredditInput(input string) string {
	if input == "" {
		return ""
	}

	s := strings.ToLower(input)
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "http://")
	s = strings.TrimPrefix(s, "www.")
	s = strings.TrimPrefix(s, "reddit.com/r/")
	s = strings.TrimPrefix(s, "old.reddit.com/r/")

	if strings.HasPrefix(s, "r/") {
		s = strings.TrimPrefix(s, "r/")
	}
	if strings.HasPrefix(s, "/r/") {
		s = strings.TrimPrefix(s, "/r/")
	}

	s = strings.TrimRight(s, "/")

	return s
}

func getRandomSubreddit() string {
	var choices []string

	for sub, weight := range config.DefaultMemeSubreddits {
		for i := 0; i < weight; i++ {
			choices = append(choices, sub)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(choices))
	return choices[idx]
}

func fetchSingleMeme(subreddit string) (*models.MemeResponse, error) {
	url := fmt.Sprintf("https://meme-api.com/gimme/%s", subreddit)

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

func GetMeme(userInput string) (*models.MemeResponse, error) {
	targetSubreddit := cleanSubredditInput(userInput)
	tryingUserRequest := (targetSubreddit != "")

	for i := 0; i < config.MemeMaxRetries; i++ {
		currentSub := targetSubreddit

		if !tryingUserRequest {
			currentSub = getRandomSubreddit()
		}

		meme, err := fetchSingleMeme(currentSub)
		if err != nil {
			if tryingUserRequest {
				tryingUserRequest = false
				continue
			}
			continue
		}

		return meme, nil
	}

	return nil, fmt.Errorf("failed to find a meme after %d retries", config.MemeMaxRetries)
}
