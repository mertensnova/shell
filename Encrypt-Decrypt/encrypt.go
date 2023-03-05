package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var files []string

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Encrypt() {
	entries, err := ioutil.ReadDir("./test")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range entries {
		files = append(files, file.Name())
	}

	for _, v := range files {
		content, err := os.ReadFile("./test/" + v)
		c, err := aes.NewCipher([]byte("this_must_be_of_32_byte_length!!"))

		if err != nil {
			log.Fatalln(err)
		}

		size := c.BlockSize()
		data := PKCS5Padding(content, size)
		txt := make([]byte, len(data))
		c.Encrypt(txt, data)

		err = os.WriteFile("./test/"+v, []byte(hex.EncodeToString(txt)), 0644)

		if err != nil {
			log.Fatalln(err)
		}

	}
}

func main() {
	start := time.Now()
	Encrypt()
	elapsed := time.Since(start)
	fmt.Printf("\n\nTime took %s", elapsed)
}
