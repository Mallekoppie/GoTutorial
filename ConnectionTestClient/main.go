package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	time "time"
)

type function func()

func MakeCall() {
	resp, err := http.Get("http://localhost:80/ConnectionTest")

	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed!")
	}
}

func MakeBetterCall() {
	req, err := http.NewRequest("GET", "http://localhost:80/ConnectionTest", nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Response error: %v", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("repsonse failed")
	}
}

var (
	httpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func MuchBetterCall() {
	//var endPoint string = "http://localhost:80/ConnectionTest"
	var endPoint string = "http://localhost:85/ConnectionTest/ConnectionTest"

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		log.Fatalf("Error Occured. %+v", err)
	}

	req.Header.Add("test", "testValue")

	// use httpClient to send request
	response, err := httpClient.Do(req)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	} else {
		// Close the connection to reuse it
		defer response.Body.Close()

		// Let's check if the work actually is done
		// We have seen inconsistencies even when we get 200 OK response
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Couldn't parse response body. %+v", err)
		}

		//log.Println("Response Body:", string(body))
	}
}

func TimeFunc(function function) {
	timer := time.Now()

	for i := 0; i < 1000; i++ {
		function()
		time.Sleep(time.Second * 1)
	}
	result := time.Since(timer)
	log.Printf("End time: %v", result.Seconds())

}

func main() {
	//TimeFunc(MakeCall)
	//TimeFunc(MakeCall)
	//TimeFunc(MakeCall)
	//TimeFunc(MakeBetterCall)
	//TimeFunc(MakeBetterCall)
	//TimeFunc(MakeBetterCall)
	TimeFunc(MuchBetterCall)
	//TimeFunc(MuchBetterCall)
	//TimeFunc(MuchBetterCall)

}
