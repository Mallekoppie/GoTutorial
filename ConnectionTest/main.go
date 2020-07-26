package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET Executed")
		log.Println("Server 2 hit")
		w.WriteHeader(http.StatusOK)
		break
	case "POST":
		fmt.Fprintf(w, "POST Executed")
		w.WriteHeader(http.StatusCreated)
		http.Error(w, "bla", 501)

		break
	default:
		w.WriteHeader(http.StatusNoContent)
		break
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path no found:", r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type KnockoutData struct {
	People []Person `json:"people"`
}

func KnockoutTest(w http.ResponseWriter, r *http.Request) {
	log.Println("KnockoutTest Hit")
	data := KnockoutData{
		People: []Person{{Name: "Test Name 1", Surname: "Test Surname 1"}, {Name: "Test Name 2", Surname: "Test Surname 2"}},
	}

	w.WriteHeader(http.StatusOK)
	json, err := json.Marshal(data)
	if err != nil {
		log.Println("Unable to marshal data: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/ConnectionTest", sayHello)
	myMux.HandleFunc("/Knockout", KnockoutTest)

	myMux.HandleFunc("/", notFound)
	fmt.Println("HTTP Server is running on port 85")
	err := http.ListenAndServe("0.0.0.0:85", myMux)
	if err != nil {
		log.Println("Unable to start: ", err.Error())
	}

}
