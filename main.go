package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"posterboi/handler"
	"time"

	"github.com/gorilla/mux"
)

func tick(ticker *time.Ticker, q *[]interface{}) {
	for t := range ticker.C {
		log.Println("Tick at", t)
		log.Println(q)
	}
}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	queue := loadQueue()

	go tick(ticker, &queue)

	r := mux.NewRouter()

	r.HandleFunc("/reddit", handler.RedditHandler(&queue)).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func loadQueue() (queue []interface{}) {
	b, err := ioutil.ReadFile("queue.json")

	if err != nil {
		queue = make([]interface{}, 0)
		return
	}

	err = json.Unmarshal(b, &queue)

	if err != nil {
		panic(err)
	}

	return
}
