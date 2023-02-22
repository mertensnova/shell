package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeTable() {
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

	m1 := make(map[string]string)
	var s []string

	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// Get body of the table
		table.Find("tbody").Each(func(i int, body *goquery.Selection) {
			body.Find("tr").Each(func(i int, rowbody *goquery.Selection) {
				rowbody.Find("td").Each(func(bi int, titlebody *goquery.Selection) {
					// Get head of the table
					table.Find("thead").Each(func(i int, head *goquery.Selection) {
						head.Find("tr").Each(func(i int, rowhead *goquery.Selection) {
							rowhead.Find("th").Each(func(i int, titlehead *goquery.Selection) {
								// Append the head to the array
								s = append(s, titlehead.Text())
							})
						})
					})

					// Save the body and head inside a map
					m1[s[bi]] = titlebody.Text()
				})

				// Convert map to JSON
				data, err := json.Marshal(m1)
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				}
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				}
				// Write inside the file
				f, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				if _, err = f.Write(data); err != nil {
					panic(err)
				}
			})

		})

	})
}

func main() {
	start := time.Now()
	ScrapeTable()
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
