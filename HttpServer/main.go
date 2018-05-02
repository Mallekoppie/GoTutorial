package main

import (
	"encoding/json"
	io "io/ioutil"
	"log"
	"net/http"
)

func readTestBody(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	test := string(body)

	log.Print(test)

}

type TestJson struct {
	Name string
}

func readJsonBody(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	object := TestJson{}

	err := json.Unmarshal(body, &object)

	if err != nil {
		log.Println("unable to unmarchal: ", err)
	}

	log.Printf("Name in body: %v", object.Name)
}

func ReturnJsonBody(w http.ResponseWriter, r *http.Request) {
	body := TestJson{Name: "Test name in response body"}

	data, err := json.Marshal(body)

	if err != nil {
		log.Println("The marshalling failed", err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	number, err2 := w.Write(data)

	if err2 != nil {
		log.Println("Something went wrong when writing response:", err)
	}

	log.Println("Body write return value: ", number)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/textBody", readTestBody)
	mux.HandleFunc("/jsonBody", readJsonBody)
	mux.HandleFunc("/returnjsonBody", ReturnJsonBody)
	log.Println("Http Server listening on port 9000 and path /test")

	http.ListenAndServe(":9000", mux)
}
