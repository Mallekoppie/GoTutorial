package main

import (
	b64 "encoding/base64"
	io "io/ioutil"
	"log"
)

func main() {

	bytes, err := io.ReadFile("jsonContent.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	data := string(bytes)

	log.Print(data)

	encodedData := b64.StdEncoding.EncodeToString(bytes)

	log.Print(encodedData)
}
