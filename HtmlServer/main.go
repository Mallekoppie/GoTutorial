package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed dist/*
var dist embed.FS

func main() {
	stripped, err := fs.Sub(dist, "dist")
	if err != nil {
		fmt.Println("Error stripping forntend")
	}
	fileServer := http.FileServer(http.FS(stripped))

	r := mux.NewRouter()
	// It's important that this is before your catch-all route ("/")
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/users", GetUsersHandler).Methods("GET")
	// Serve static assets directly.
	//r.PathPrefix("/assets").Handler(fileServer)
	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").Handler(fileServer)

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "127.0.0.1:9999",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"id":        "123",
		"timeStamp": time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write(b)
}
