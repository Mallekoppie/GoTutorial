package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func callConnectionTestService(writer http.ResponseWriter, reaquest *http.Request) {

	remoteServiceName := os.Getenv("remoteservice")

	if len(remoteServiceName) < 1 {
		log.Fatalln("Unable to retrieve remoteService environment variable")
		return
	} else {
		log.Println("Remote Service Name: " + remoteServiceName)
	}

	resp, err := http.Get("http://" + remoteServiceName + ":80/ConnectionTest")

	if err != nil {
		log.Fatalln("Error calling connection test service: " + err.Error())
	}

	if resp != nil {

		data, readErr := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if readErr != nil {
			log.Fatalln("Error reading body: " + readErr.Error())
		}

		log.Println("Connection test call successfull: " + string(data))
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
	} else {
		log.Fatalln("Response object is nil")
	}

}

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/call", callConnectionTestService)

	http.ListenAndServe("0.0.0.0:80", myMux)
}
