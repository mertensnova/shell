package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var files []string

func Decrypt() {
	entries, err := ioutil.ReadDir("./test")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range entries {
		files = append(files, file.Name())
	}

	for _, v := range files {
		content, err := os.ReadFile("./test/" + v)
		txt, _ := hex.DecodeString(string(content))
		c, err := aes.NewCipher([]byte("this_must_be_of_32_byte_length!!"))
		if err != nil {
			log.Fatalln(err)
		}

		data := make([]byte, len(txt))
		c.Decrypt(data,[]byte(txt))

		err = os.WriteFile("./test/" + v,data,0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	start := time.Now()
	Decrypt()
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
