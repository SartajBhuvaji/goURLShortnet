package main

import (
	"log"
	"net/http"

	"github.com/SartajBhuvaji/api"
)

func main() {
	http.HandleFunc("/shorten", api.ShortenURLHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
