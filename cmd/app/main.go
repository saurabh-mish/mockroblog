package main

import (
	"log"
	"net/http"
	"time"

	"mockroblog/pkg/routes"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Serve)

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
