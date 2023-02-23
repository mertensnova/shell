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

// A Writer wraps a csv.Writer to write records as maps of column names to
// values, instead of lists of values.
type Writer struct {
	w      *csv.Writer
	header []string
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w *csv.Writer, header []string) *Writer {
	return &Writer{
		w:      w,
		header: header,
	}
}

// WriteHeader writes the CSV header to w along with any necessary quoting.
func (w *Writer) WriteHeader() error {
	return w.w.Write(w.header)
}

// Writer writes a single CSV record to w along with any necessary quoting. A
// record is a map of column names to values. Only columns present in the
// Writer's header are written, and in the order they appear in the header.
func (w *Writer) Write(record map[string]string) error {
	s := make([]string, len(w.header))
	for i, name := range w.header {
		s[i] = record[name]
	}
	return w.w.Write(s)
}

// WriteAll writes multiple CSV records to w using Write and then calls w.Flush.
func (w *Writer) WriteAll(records []map[string]string) error {
	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			return err
		}
	}
	w.w.Flush()
	return w.w.Error()
}

func WriterToCSV(data []map[string]string, header []string) {
	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := NewWriter(csv.NewWriter(f), header)

	err = w.WriteHeader()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(header)
	// for i := 0; i < len(data); i++ {
		// err := w.Write(data[i])
		for _, v := range data {
			fmt.Println(v)
		}
		
	// }

	f.Close()
}

func WriteToJSON(data []map[string]string) {
	// Write inside the file
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
	var data_array []map[string]string
	var header []string

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
					data_map[header[index]] = titlebody.Text()
				})
				//  Append the map in an array
				data_array = append(data_array, data_map)
			})
		})

	})
	// fmt.Println(header)

	WriterToCSV(data_array, header)
	// WriteToJSON(data_array)

}

func main() {
	start := time.Now()
	ScrapeTable()
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
