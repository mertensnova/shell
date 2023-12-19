package gofuncs

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ReadFile(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func SendRequest(url string, method string, payload interface{}) string {

	client := &http.Client{}

	request, err := http.NewRequest(method, url, nil)
	request.Header.Set("User-Agent", RandomUserAgents())

	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("Error reading response body:", err)
	}

	return string(body)
}

func RandomUserAgents(filename string) string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("File is empty")
	}

	r := rand.Intn(len(lines))

	return lines[r]
}
