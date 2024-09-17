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

type job struct {
	id        string
	title     string
	company   string
	location  string
	condition []string
	jobDate   string
	badge     string
}

func main() {
	totalJobs := []job{}
	jobsCount := getJobsCount()
	totalPages := (jobsCount / defaultPerPage) + 1

	for i := 0; i < totalPages; i++ {
		jobs := getPage(i + 1)
		totalJobs = append(totalJobs, jobs...)
	}

	fmt.Println(totalJobs)
}

func extractJob(card *goquery.Selection) job {
	id, _ := card.Attr("value")
	id = cleanString(id)
	title := cleanString(card.Find("div.area_job").Find("h2.job_tit").Find("a").Find("span").Text())
	company := cleanString(card.Find("div.area_corp").Find("strong.corp_name").Text())
	location := ""
	condition := []string{}
	conditionResult := card.Find("div.area_job").Find("div.job_condition").Find("span")
	conditionResult.Each(func(i int, cd *goquery.Selection) {
		conditionAtag := cd.Find("a")
		if conditionAtag.Length() > 0 {
			conditionAtag.Each(func(i int, conditionLocation *goquery.Selection) {
				if location != "" {
					location += " "
				}
				location += cleanString(conditionLocation.Text())
			})
		} else {
			condition = append(condition, cleanString(cd.Text()))
		}
	})
	jobDate := cleanString(card.Find("div.area_job").Find("div.job_date").Find("span.date").Text())
	badge := cleanString(card.Find("div.area_badge").Find("span").Text())

	return job{
		id:        id,
		title:     title,
		company:   company,
		location:  location,
		condition: condition,
		jobDate:   jobDate,
		badge:     badge,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPage(page int) []job {
	jobs := []job{}
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	response, err := http.Get(pageURL)
	checkError(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	searchResult := doc.Find("div.item_recruit")
	searchResult.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
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
