package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Authorizer interface {
	Add(req *http.Request)
}

type Client struct {
	Authorizer Authorizer
	Client     *http.Client
	Host       string //https://api.twitter.com/
}

type CreateTweetRequest struct {
	Tweet string `json:"text,omitempty"`
}

type HTTPError struct {
	Status     string
	StatusCode int
	URL        string
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("twitter [%s] status: %s code: %d", h.URL, h.Status, h.StatusCode)
}

type Error struct {
	Parameters interface{} `json:"parameters"`
	Message    string      `json:"message"`
}

func (c *Client) CreateTweet(ctx context.Context, tweet CreateTweetRequest) (string, error) {
	body, err := json.Marshal(tweet)
	if err != nil {
		return "", fmt.Errorf("create tweet request: %w", err)
	}

	endpoint := tweetCreateEndpoint.url(c.Host)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("create tweet request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	c.Authorizer.Add(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("create tweet response: %w", err)
	}

	defer resp.Body.Close()

	return "", err
}
