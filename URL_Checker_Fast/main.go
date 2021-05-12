package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {

	c := make(chan requestResult)
	results := make(map[string]string)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://github.com",
		"https://amazon.com",
		"https://www.reddit.com",
		"https://www.soundcloud.com",
		"https://www.linkedon.com", //purposefully incorrect to test error codes
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(status, url)
	}

}

//hit URL, only sends through channel
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "SUCCESS"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED "
	}

	c <- requestResult{url: url, status: status}
}
