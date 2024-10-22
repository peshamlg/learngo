package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type job struct {
	id        string
	title     string
	company   string
	location  string
	condition []string
	jobDate   string
	badge     string
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

// CleanString is a function that cleans the string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getJobsCount(url string) int {
	jobsCount := 0
	response, err := http.Get(url)
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

func getPage(page int, url string, mc chan<- []job) {
	jobs := []job{}
	c := make(chan job)
	pageURL := url + "&recruitPage=" + strconv.Itoa(page)
	response, err := http.Get(pageURL)
	checkError(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	searchResult := doc.Find("div.item_recruit")
	searchResult.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchResult.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mc <- jobs
}

func extractJob(card *goquery.Selection, c chan<- job) {
	id, _ := card.Attr("value")
	id = CleanString(id)
	title := CleanString(card.Find("div.area_job").Find("h2.job_tit").Find("a").Find("span").Text())
	company := CleanString(card.Find("div.area_corp").Find("strong.corp_name").Text())
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
				location += CleanString(conditionLocation.Text())
			})
		} else {
			condition = append(condition, CleanString(cd.Text()))
		}
	})
	jobDate := CleanString(card.Find("div.area_job").Find("div.job_date").Find("span.date").Text())
	badge := CleanString(card.Find("div.area_badge").Find("span").Text())

	c <- job{
		id:        id,
		title:     title,
		company:   company,
		location:  location,
		condition: condition,
		jobDate:   jobDate,
		badge:     badge,
	}
}

func writeJobs(jobs []job) {
	file, err := os.Create("jobs.csv")
	checkError(err)

	writer := csv.NewWriter(file)
	defer writer.Flush()
	defer file.Close() // main.go의 defer os.Remove("jobs.csv")가 동작하려면 여기서 닫아야 함

	headers := []string{"ID", "Title", "Company", "Location", "Condition", "Job Date", "Badge", "Apply URL"}
	writer.Write(headers)

	for _, job := range jobs {
		jobSlice := []string{
			job.id,
			job.title,
			job.company,
			job.location,
			strings.Join(job.condition, " "),
			job.jobDate,
			job.badge,
			"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id,
		}
		jwErr := writer.Write(jobSlice)
		checkError(jwErr)
	}
}

// Scrape is a function that scrapes the jobs from the saramin website
func Scrape(term string) {
	totalJobs := []job{}
	url := "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	defaultPerPage := 40
	jobsCount := getJobsCount(url)
	totalPages := (jobsCount / defaultPerPage) + 1
	mc := make(chan []job)

	for i := 0; i < totalPages; i++ {
		go getPage(i+1, url, mc)
	}

	for i := 0; i < totalPages; i++ {
		jobs := <-mc
		totalJobs = append(totalJobs, jobs...)
	}

	writeJobs(totalJobs)
	fmt.Println("Done! Total jobs:", len(totalJobs))
}
