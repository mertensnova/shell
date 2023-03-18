package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"
)

func JPEGHandler(path string, f *os.File, res *http.Response) {
	myImage, err := jpeg.Decode(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	if err = jpeg.Encode(f, myImage, nil); err != nil {
		log.Printf("failed to encode: %v", err)
	}
	log.Println(filepath.Base(path) + " has been downloaded")

}
func PNGHandler(path string, f *os.File, res *http.Response) {

	myImage, err := png.Decode(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	if err = png.Encode(f, myImage); err != nil {
		log.Printf("failed to encode: %v", err)
	}

	log.Println(filepath.Base(path) + " has been downloaded")
}

func DownloadImage(url string, path string) {
	res, err := http.Get(url + path)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	defer res.Body.Close()

	if _, err := os.Stat("images"); os.IsNotExist(err) {
		if err := os.Mkdir("images", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create("./images/" + filepath.Base(path))

	if err != nil {
		log.Fatalln(err)
	}

	switch filepath.Ext(path) {
	case ".png":
		PNGHandler(path, f, res)
	case ".jpg":
		JPEGHandler(path, f, res)
	}
}

func main() {
	start := time.Now()

	url := flag.String("u", "https://www.google.hu/", "URL page you want to download")
	flag.Parse()
	png_pattern := regexp.MustCompile(`[A-Z0-9a-z\.\/_-]*(\.png)`)
	jpeg_pattern := regexp.MustCompile(`[A-Z0-9a-z\.\/_-]*(\.jpeg)`)

	res, err := http.Get(*url)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if png_pattern.MatchString(string(bytes[:])) {
		DownloadImage(*url, path.Clean(png_pattern.FindString(string(bytes[:]))))
	} else if jpeg_pattern.MatchString(string(bytes[:])) {
		DownloadImage(*url, path.Clean(jpeg_pattern.FindString(string(bytes[:]))))
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
