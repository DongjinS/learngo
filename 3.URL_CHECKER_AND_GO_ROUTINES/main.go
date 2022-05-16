package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main () {
	urlCHK()
	// goroutinePrac()
}

func goroutinePrac(){
	c := make(chan string)
	people := [6]string {"DJ","SE","SH","HM","HJ","JW"}
	for _, person := range people{
		go isSexy(person, c)
	}
	for i:=0; i<len(people); i++ {
		fmt.Println("Waiting for", i)
		fmt.Println("Received this message:",<-c)
	}
}

func isSexy(name string, c chan string){
	c<-name+" is sexy"
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

type requestResult struct {
	url string;
	status string;
}

func urlCHK() {
	results := make(map[string]string)
	// results := map[string]string{}

	ch := make(chan requestResult)

	urls := []string {
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for i:=0;i<len(urls);i++{
		result := <-ch
		results[result.url] = result.status
	}

	for url, status := range results{
		fmt.Println(url, status)
	}

}

func hitURL(url string, ch chan<- requestResult) {
	status := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		status = "FAILED"
	}
	ch<-requestResult{url:url, status: status}
}
