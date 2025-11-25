package models

type MemeResponse struct {
	PostLink  string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
	NSFW      bool     `json:"nsfw"`
	Author    string   `json:"author"`
	Ups       int      `json:"ups"`
	Preview   []string `json:"preview"`
}
