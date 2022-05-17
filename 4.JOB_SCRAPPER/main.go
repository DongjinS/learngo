package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

type extractedJob struct {
	id string;
	title string;
	company string;
	location string;
	salary string;
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	ch := make(chan []extractedJob)
	var jobs []extractedJob
	totalPages := getPages()

	for i:=0;i<totalPages;i++{
		go getPage(i, ch)
	}

	for i:=0;i<totalPages;i++{
		jobs = append(jobs, <-ch...)
	}

	writeJobs(jobs)
	writeJobswithGoRoutines(jobs)

	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob){
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Company", "Location", "Salary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs{
		jobSlice := []string{"https://kr.indeed.com/jobs?q=python&l&vjk="+job.id, job.title, job.company, job.location, job.salary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func writeJobswithGoRoutines(jobs []extractedJob) {
	csv, err := ccsv.NewCsvWriter("sample.csv")
	checkErr(err)
	
	defer csv.Close()
	
	csv.Write([]string{"Link", "Title", "Location", "Salary"})
	done := make(chan bool)
	for _, job := range jobs {
		go func(job extractedJob) {
			csv.Write([]string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary})
			done <- true
		}(job)
	}
	for i := 0; i < len(jobs); i++ {
		<-done
	}
}

func getPage(page int, ch chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesing:",pageURL)
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	cnt := 0
	doc.Find(".cardOutline").Each(func(i int, s *goquery.Selection){
		go extractJob(s,c)
		cnt +=1
	})
	for i:=0;i<cnt;i++{
		jobs = append(jobs,<-c)
	}

	ch<-jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	resultContent := card.Find(".resultContent") 
	atag := resultContent.Find("a")
	id, exists := atag.Attr("data-jk")
	checkExist(exists)

	title, exists := atag.Find("span").Attr("title")
	checkExist(exists)
	
	companyInfo := resultContent.Find(".companyInfo")
	name := companyInfo.Find("span").Text()
	location := companyInfo.Find(".companyLocation").Text()
	salary := resultContent.Find(".salaryOnly").Text()
	if salary==""{
		salary = "??"
	} 
	c<-extractedJob{
		id: id,
		title: title, 
		company: name, 
		location: location, 
		salary: salary,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	
	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})
	return pages
}

func checkExist(exists bool){
	if !exists{
		log.Fatalln("not existed in doc")
	}
}

func checkErr(err error){
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response){
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}