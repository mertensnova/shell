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

func WriteToCSV(header []string) {
	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := NewWriter(csv.NewWriter(f), header)

	err = w.WriteHeader()
	if err != nil {
		log.Fatal(err)
	}

	// err = w.Write(header)

	if err != nil {
		log.Fatalln(err)
	}

	f.Close()
}

func main() {
	start := time.Now()
	var headers []string

	var rows []string
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
			// rows = append(rows, s.Text())
		})
	})
	var data map[string][]string

	for _, header := range headers {
		for _, row := range rows {
			data[header] = append(data[header], row)
			// data = append(data, map[string]string{
			// 	header: row,
			// })
			fmt.Println(data)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
