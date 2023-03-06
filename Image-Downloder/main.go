package main

import (
	"bufio"
	"fmt"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func DownloadImage(url string, n string) {
	res, err := http.Get("https://www.pexels.com/search/png" + url)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	defer res.Body.Close()

	f, err := os.Create("./images/" + n + ".png")
	if err != nil {
		panic(err)
	}
	// Keep an in memory copy.
	myImage, err := png.Decode(res.Body)

	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, myImage); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}

func main() {

	url := "https://cs50.readthedocs.io/ide/online/#working-with-files"
	re := regexp.MustCompile(`[A-Z0-9a-z\.\/_]*(\.png)`)

	out, err := os.Create("file.txt")

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	defer res.Body.Close()

	io.Copy(out, res.Body)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++
		if re.MatchString(scanner.Text()) {
			fmt.Println(strings.Split(re.FindString(scanner.Text()), "../../"))
			// strings.Split(re.FindString(scanner.Text()), "../../")
			// DownloadImage(re.FindString(scanner.Text()), strconv.Itoa(count))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer out.Close()
}
