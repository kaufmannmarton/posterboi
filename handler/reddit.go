package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"posterboi/model"
)

func RedditHandler(queue *[]interface{}) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var rr model.RedditRequest

		json.Unmarshal(body, &rr)

		if err = rr.Validate(); err != nil {
			b, _ := json.Marshal(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(b))
			return
		}

		for _, subreddit := range rr.Subreddits {
			*queue = append(*queue, model.RedditPost{
				ClientID:     rr.ClientID,
				ClientSecret: rr.ClientSecret,
				Username:     rr.Username,
				Password:     rr.Password,
				UserAgent:    rr.UserAgent,
				Subreddit:    subreddit,
				PostTitle:    rr.PostTitle,
				PostURL:      rr.PostURL,
				Instance:     rr.Instance,
			})
		}

		w.WriteHeader(http.StatusOK)
	})
}
