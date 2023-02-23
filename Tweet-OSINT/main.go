package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request the HTML page.
	username := os.Args[1]

	res, err := http.Get("https://twitter.com/" + username)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	// doc.Find("main.ss-1dbjc4n").Find("div").Each(func(i int, s *goquery.Selection) {
		fmt.Println(doc.Find("main.ss-1dbjc4n").Size())
	// })	
	// .Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Find("h2").Text())
	// })

	if err != nil {
		log.Println(err)
	}
}
