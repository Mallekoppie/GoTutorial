package main

import (
	b64 "encoding/base64"
	json "encoding/json"
	io "io/ioutil"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

type Test struct {
	Name   string `json:"name,omitempty"`
	Method string `json:"method,omitempty"`
	Body   string `json:"body,omitempty"`
}

type Config struct {
	Name      string
	Nicknames []string
	Tests     []Test
}

func GetConfig() Config {
	configuration := Config{}
	err := gonfig.GetConf("config.json", &configuration)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print(configuration)

	data, err := b64.StdEncoding.DecodeString(configuration.Tests[0].Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print(string(data))

	if len(configuration.Tests[1].Body) < 1 {
		log.Print("nilTest body is nil")
	}

	return configuration
}

func testGetConfig() {
	configuration := GetConfig()

	log.Print(configuration.Name)
}

func testWriteConfig() {
	config := Config{Name: "Write", Nicknames: []string{"test", "test2"}}

	data, err := json.Marshal(config)

	if err != nil {
		log.Println("Failed to marshall config")
	}

	io.WriteFile("writtenconfig.json", data, os.ModeExclusive)

}

func main() {
	//testGetConfig()

	testWriteConfig()

}
