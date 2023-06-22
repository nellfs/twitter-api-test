package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	bearer := os.Getenv("TWITTER_BEARER_TOKEN")

	token := flag.String("token", bearer, "twitter API token")
	text := flag.String("text", "esse tweet foi enviado pela api do twitter usando go :)", "twitter text")
	flag.Parse()

	client := &Client{
		Authorizer: authorize{
			Token: *token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com/",
	}
	req := CreateTweetRequest{Tweet: *text}
	fmt.Println("Callout to create tweet callout")

	tweetResponse, err := client.CreateTweet(context.Background(), req)
	if err != nil {
		log.Panicf("create tweet error: %v", err)
	}

	enc, err := json.MarshalIndent(tweetResponse, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}
