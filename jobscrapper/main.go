package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"
var defaultPerPage int = 40

func main() {
	jobsCount := getJobsCount()
	totalPages := (jobsCount / defaultPerPage) + 1

	for i := 0; i < totalPages; i++ {
		getPage(i + 1)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
}

func getJobsCount() int {
	jobsCount := 0
	response, err := http.Get(baseURL)
	checkError(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	pageJobsCount := doc.Find("div.header").Find("span.cnt_result").Text()
	re := regexp.MustCompile(`\d+`)
	numStr := strings.Join(re.FindAllString(pageJobsCount, -1), "")
	count, err := strconv.Atoi(numStr)
	checkError(err)
	jobsCount += count

	return jobsCount
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(response *http.Response) {
	if response.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", response.StatusCode)
	}
}
