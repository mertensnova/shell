package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {

	url := "https://socketloop.com"
	re := regexp.MustCompile(`(http:)[A-Z0-9a-z\.\/_]*(\.mp4)`)

	fmt.Println(re.FindString(`<img id="clip-4182439" data-hls-src="http://mpegmedia.abc.net.au/abcforkids/video/boj/boj_SneezySnufferoos_MusicVid03_HD.mp4" src="http://www.abc.net.au/reslib/201502/r1390108_19759857.jpg" alt="">`))

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	htmlData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	videoMatchSlice := re.FindAllStringSubmatch(string(htmlData), -1)
	for _, item := range videoMatchSlice {
		log.Println("Video found : ", item)
	}
}

