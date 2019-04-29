package model

type RedditPost struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	UserAgent    string `json:"user_agent"`
	Subreddit    string `json:"subreddit"`
	PostTitle    string `json:"post_title"`
	PostURL      string `json:"post_url"`
	Instance     string `json:"instance"`
}
