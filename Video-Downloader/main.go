package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// Request the HTML page.
	res, err := http.Get("https://www.ted.com/talks/johan_rockstrom_10_years_to_transform_the_future_of_humanity_or_destabilize_the_planet")
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

	fmt.Println(doc)

	fmt.Println("Hello World")
}
