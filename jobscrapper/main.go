package main

import (
	"jobscrapper/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrpae(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrpae)
	e.Logger.Fatal(e.Start(":1323"))
}
