package main

import (
	io "io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://novelonlinefree.info/chapter/battle_through_the_heavens/chapter_1042")
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error downloading: ", err)
	}

	data, _ := io.ReadAll(resp.Body)

	err2 := io.WriteFile("testpage.html", data, os.ModeExclusive)

	if err2 != nil {
		log.Println("Error writing to disk: ", err2)
	}

}
