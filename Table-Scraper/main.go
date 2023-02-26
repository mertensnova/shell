package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	start := time.Now()
	var headers []string

	rows := []string{}
	// var data_arr []map[string]string
	// Request the HTML page.
	res, err := http.Get("https://datatables.net/examples/styling/display.html")
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

	doc.Find("table").Find("thead").Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("th").Each(func(i int, s *goquery.Selection) {
			headers = append(headers, s.Text())
		})
	})

	doc.Find("table").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			rows = append(rows, s.Text())
		})
	})
	
	data_map := make(map[string]string, 7)
	for i, row := range rows {

		data_map[headers[i%len(headers)]] = row

		fmt.Println(data_map)
		
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
