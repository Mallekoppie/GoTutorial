package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello hit")
	w.Header().Add("Location", "http://localhost:8180/auth/realms/SpringBootKeycloak/protocol/openid-connect/auth?client_id=redirectapp&scope=openid&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A10003%2Fauth&state=123123123")
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func auth(w http.ResponseWriter, r *http.Request) {
	log.Println("Auth hit")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Println("Error reading body: " + err.Error())
	}

	log.Println(string(body))

	for header := range r.Header {
		log.Printf("%v %v \n", header, r.Header.Get(header))
	}

	for item := range r.Form {
		log.Println("Printing form values")
		log.Printf("%v %v \n", item, r.Form.Get(item))
		log.Printf("%v %v \n", r.FormValue(item))
	}

	log.Println("RequestURI: " + r.RequestURI)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/auth", auth)

	http.ListenAndServe(":10003", mux)
}
