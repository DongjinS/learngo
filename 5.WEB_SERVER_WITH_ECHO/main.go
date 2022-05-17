package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DongjinS/learngo/5.WEB_SERVER_WITH_ECHO/scrapper"
	"github.com/labstack/echo"
)


func handleHome(c echo.Context) error {
	return c.File("home.html")
} 
	
func handleScrape(c echo.Context) error {
	defer cleanUp()
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	fmt.Println(term)
	scrapper.Scrape(term)
	c.Attachment("jobs.csv", "jobs.csv")
	return nil
}

func cleanUp() {
	os.Remove("jobs.csv")
	os.Remove("sample.csv")
}

func main () {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}