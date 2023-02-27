package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	start := time.Now()

	// Request the HTML page.
	res, err := http.Get("https://www.the-numbers.com/weekend-box-office-chart")
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	size := doc.Find("table").Find("tbody").Find("tr").Size()

	records := make([][]string, size+1)

	doc.Find("table").Find("thead").Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("th").Each(func(i int, s *goquery.Selection) {
			records[0] = append(records[0], s.Text())
		})
	})

	doc.Find("table").Find("tbody").Find("tr").Each(func(ix int, s *goquery.Selection) {
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			records[ix+1] = append(records[ix+1], s.Text())
		})
	})

	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(records)

	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
