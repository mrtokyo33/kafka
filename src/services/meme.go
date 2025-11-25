package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"kafka/src/config"
	"kafka/src/models"
)

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

func GetMeme(userSubreddit string, forceNSFW bool) (*models.MemeResponse, error) {
	targetSubreddit := userSubreddit

	for i := 0; i < config.MemeMaxRetries; i++ {

		if userSubreddit == "" {
			targetSubreddit = getRandomSubreddit()
		}

		meme, err := fetchSingleMeme(targetSubreddit)
		if err != nil {
			return nil, err
		}

		if forceNSFW {
			if meme.NSFW {
				return meme, nil
			}
			continue
		}

		return meme, nil
	}

	return nil, fmt.Errorf("failed to find a meme matching criteria after retries")
}
