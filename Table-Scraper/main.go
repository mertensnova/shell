package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Data map[string]interface{}

var myMapSlice []Data

func WriteToCSV(data map[string]string, header []string) {
	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	data_array := []map[string]string{data}
	unique_headers := RemoveDuplicates(header)
	w := NewWriter(csv.NewWriter(f), unique_headers)

	err = w.WriteHeader()
	if err != nil {
		log.Fatal(err)
	}

	err = w.WriteAll(data_array)

	if err != nil {
		log.Fatalln(err)
	}

	f.Close()
}

func WriteToJSON(data map[string]string) {

	byte, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	f, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err = f.Write(byte); err != nil {
		log.Println(err)
	}

}

func ScrapeTable() {
	data_map := make(map[string]string)
	var header []string
	// var bod []string

	// Request the HTML page.
	res, err := http.Get("https://www.the-numbers.com/box-office-records/worldwide/all-movies/cumulative/all-time")
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

	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// Get body of the table
		table.Find("tbody").Each(func(i int, body *goquery.Selection) {
			body.Find("tr").Each(func(i int, rowbody *goquery.Selection) {
				rowbody.Find("td").Each(func(index int, titlebody *goquery.Selection) {
					// Get head of the table
					table.Find("thead").Each(func(i int, head *goquery.Selection) {
						head.Find("tr").Each(func(i int, rowhead *goquery.Selection) {
							rowhead.Find("th").Each(func(i int, titlehead *goquery.Selection) {
								// Append the head to the array
								header = append(header, titlehead.Text())
							})
						})
					})
					// Save the body and head inside a map

					data_map[header[index]] = fmt.Sprintf(titlebody.Text())
				})

				WriteToJSON(data_map)
				WriteToCSV(data_map, header)
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
