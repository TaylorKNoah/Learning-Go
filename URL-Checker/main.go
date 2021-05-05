package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	results          = make(map[string]string)
	errRequestFailed = errors.New("request failed")
	errStatusCode    = errors.New("status code above 400")
)

func main() {
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
		result := "OK"
		err := hitURL(url)

		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

//hit URL
func hitURL(url string) error {

	//Output to user current status
	fmt.Println("Checking", url)

	//response, error
	resp, err := http.Get(url)

	//400's where things go wrong
	if err != nil {
		fmt.Println("ERROR:", err)
		return errRequestFailed
	} else if resp.StatusCode >= 400 {
		fmt.Println("STATUS CODE:", resp.StatusCode)
		return errStatusCode
	}

	return nil
}
