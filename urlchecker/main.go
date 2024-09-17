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
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.youtube.com",
		"https://www.naver.com",
		"https://www.daum.net",
		"https://www.nate.com",
		"https://www.kakao.com",
		"https://www.line.com",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	response, err := http.Get(url)
	status := "Success"
	if err != nil || response.StatusCode >= 400 {
		status = "Failed"
	}
	c <- requestResult{url: url, status: status}
}
