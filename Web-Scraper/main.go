package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var m1 map[string]string

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
	m1 := make(map[string]string)
	doc.Find("table").Each(func(i int, table *goquery.Selection) {

		table.Find("tbody").Each(func(i int, body *goquery.Selection) {
			body.Find("tr").Each(func(i int, rowbody *goquery.Selection) {
				rowbody.Find("td").Each(func(i int, titlebody *goquery.Selection) {
					table.Find("thead").Each(func(i int, head *goquery.Selection) {
						head.Find("tr").Each(func(i int, rowhead *goquery.Selection) {
							rowhead.Find("th").Each(func(i int, titlehead *goquery.Selection) {
								fmt.Println(titlehead.Text(), titlebody.Text())
							})
						})
					})
				})

			})
		})

	})

	for key, ele := range m1 {
		fmt.Println(key, ele)
	}
}

func main() {
	start := time.Now()
	ExampleScrape()
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}

// jsonStr, err := json.Marshal(m1)
// if err != nil {
// 	fmt.Printf("Error: %s", err.Error())
// } else {
// 	fmt.Println(string(jsonStr))
// }

// file, _ := json.MarshalIndent(string(jsonStr), "", " ")

// _ = ioutil.WriteFile("test.json", file, 0644)

// table.Find("thead").Each(func(i int, head *goquery.Selection) {
// 	head.Find("tr").Each(func(i int, rowhead *goquery.Selection) {
// 		rowhead.Find("th").Each(func(i int, s *goquery.Selection) {
// 			m1[s.Text()] = b.Text()
// 		})
// 	})
// })
