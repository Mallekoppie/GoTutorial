package main

import (
	"fmt"
	"net/http"
)

func hello(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Hello"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	//err := http.ListenAndServeTLS("0.0.0.0:10000", "G:/temp/localhost.cer", "G:/temp/localhost.pkcs8", mux)
	err := http.ListenAndServeTLS("0.0.0.0:10000", "/localhost.cer", "/localhost.pkcs8", mux)

	if err != nil {
		fmt.Println("Error hosting:" + err.Error())
	}

}
