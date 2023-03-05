package main

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"flag"
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

func GenerateKey() string {
	f, err := os.Create("key.key")
	key := make([]byte, 16)

	_, err = rand.Read(key)
	if err != nil {
		log.Fatalln(err)
	}
	f.WriteString(hex.EncodeToString(key))

	return hex.EncodeToString(key)
}

func Encrypt(folder string, key string) {
	entries, err := ioutil.ReadDir(folder)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range entries {
		files = append(files, file.Name())
	}

	for _, v := range files {
		content, err := os.ReadFile("./test/" + v)
		c, err := aes.NewCipher([]byte(key))

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

	folder_name := flag.String("path", "./test", "The folder you want to encrypt")

	key := GenerateKey()
	Encrypt(*folder_name, key)

	elapsed := time.Since(start)

	fmt.Printf("\n\nTime took %s", elapsed)
}
