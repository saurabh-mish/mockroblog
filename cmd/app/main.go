package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := &http.Server{
		Addr: "localhost:8080",
		Handler: mux,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}


func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}
