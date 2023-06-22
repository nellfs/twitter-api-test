package main

import "fmt"

type endpoint string

const tweetCreateEndpoint endpoint = "2/tweets"

func (e endpoint) url(host string) string {
	return fmt.Sprintf("%s%s", host, string(e))
}
