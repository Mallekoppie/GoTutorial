package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET Executed")
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

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/ConnectionTest", sayHello)
	myMux.HandleFunc("/", notFound)
	fmt.Println("HTTP Server is running")
	http.ListenAndServe(":80", myMux)

}
