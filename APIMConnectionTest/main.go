package main

import (
	"encoding/json"
	io "io/ioutil"
	"log"
	"net/http"
)

func hello(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Hello world"))
		break
	default:
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Incorrect method"))
		break
	}
}

type TestJson struct {
	Name string
}

func onlyPost(resp http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodPost:
		body, _ := io.ReadAll(req.Body)
		defer req.Body.Close()

		object := TestJson{}

		err := json.Unmarshal(body, &object)

		if err != nil {
			log.Println("unable to unmarchal: ", err)
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("Unable to unmarshal request"))
			return
		}

		resp.WriteHeader(http.StatusAccepted)
		resp.Write([]byte("Correct request. Name: " + object.Name))

		break
	default:
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Incorrect method type"))
		break
	}

}

func processHeaders(w http.ResponseWriter, r *http.Request) {

	testOne := r.Header.Get("test-one")
	log.Println("Recevied header one: " + testOne)

	testTwo := r.Header.Get("test-two")
	log.Println("Received header two: " + testTwo)

	w.Header().Add("Response-Test-One", "test for one")
	w.Header().Add("Response-Test-Two", "test for two")
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/onlypost", onlyPost)
	mux.HandleFunc("/headers", processHeaders)

	log.Println("Starting service on port 9000")
	err := http.ListenAndServe("0.0.0.0:9000", mux)

	if err != nil {
		log.Fatalln("Error serving requests: " + err.Error())
	}
}
