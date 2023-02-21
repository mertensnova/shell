package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// var m map[string]string

func ExampleScrape() {
  // Request the HTML page.
  res, err := http.Get("https://www.the-numbers.com/box-office-records/worldwide/all-movies/cumulative/all-time")
  if err != nil {
    log.Fatal(err)
  }

  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  
  // Find the review items
//   m1 := make(map[string]string)  
  doc.Find("table").Each(func(i int, s *goquery.Selection) {
	// var h,d string
	s.Find("thead").Each(func(i int, head *goquery.Selection) {
		  fmt.Println(head.Find("th").Text())

	})
})


}

func main() {
  ExampleScrape()
}

// s.Find("tbody").Each(func(i int, de *goquery.Selection) {
	// 	// fmt.Println(de.Find("tr").Text())
	// 	d = de.Find("tr").Text()
	// })

	// m1[h] = d