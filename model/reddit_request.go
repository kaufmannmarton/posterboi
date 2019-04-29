package model

import (
	"math"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RedditRequest struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	UserAgent    string   `json:"user_agent"`
	Subreddits   []string `json:"subreddits"`
	PostTitle    string   `json:"post_title"`
	PostURL      string   `json:"post_url"`
	Instance     string   `json:"instance"`
}

func (rr RedditRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.ClientID, validation.Required),
		validation.Field(&rr.ClientSecret, validation.Required),
		validation.Field(&rr.Username, validation.Required),
		validation.Field(&rr.Password, validation.Required),
		validation.Field(&rr.UserAgent, validation.Required),
		validation.Field(
			&rr.Subreddits,
			validation.Required,
			validation.Length(0, math.MaxInt8),
		),
		validation.Field(&rr.PostTitle, validation.Required),
		validation.Field(&rr.PostURL, validation.Required),
		validation.Field(&rr.Instance, validation.Required, is.URL),
	)
}
