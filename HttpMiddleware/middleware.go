package main

import (
	"log"
	"net/http"
)

func Middleware1(innerHandler http.Handler, someParam string) http.Handler {
	log.Println("Setup Middleware1 with param: ", someParam)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before Executing Middleware1 with param: ", someParam)
		innerHandler.ServeHTTP(w, r)
		log.Println("After Executing Middleware1 with param: ", someParam)
	})
}

func Middleware2(innerHandler http.Handler, someParam string) http.Handler {
	log.Println("Setup Middleware2 with param: ", someParam)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before Executing Middleware2 with param: ", someParam)
		innerHandler.ServeHTTP(w, r)
		log.Println("After Executing Middleware2 with param: ", someParam)
	})
}
