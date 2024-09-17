package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	var results = map[string]string{}
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.youtube.com",
	}
	for _, url := range urls {
		result := "Success"
		err := hitURL(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400 {
		return errors.New("request failed")
	}
	return nil
}
