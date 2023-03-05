package main

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var files []string

func GetKey(key_file string) string {
	f, err := os.Open(key_file)
	var key string
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		key = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return key
}

func Decrypt(folder string, key string) {
	entries, err := ioutil.ReadDir(folder)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range entries {
		files = append(files, file.Name())
	}

	for _, v := range files {
		content, err := os.ReadFile("./test/" + v)
		txt, _ := hex.DecodeString(string(content))
		c, err := aes.NewCipher([]byte(key))
		if err != nil {
			log.Fatalln(err)
		}

		data := make([]byte, len(txt))
		c.Decrypt(data, []byte(txt))

		err = os.WriteFile("./test/"+v, data, 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	start := time.Now()

	folder_name := flag.String("path", "./test", "The folder you want to decrypt")
	key_file := flag.String("key", "key.key", "The name of the key file")
	flag.Parse()
	key := GetKey(*key_file)

	Decrypt(*folder_name, key)
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
