package main

import (
	"log"
	"net/http"

	"github.com/SartajBhuvaji/api"
	"github.com/SartajBhuvaji/utils"
)

func main() {

	// Create a new RedisClient
	redisClient, err := utils.SetupRedis()
	if err != nil {
		return
	}
	defer redisClient.Close() // Close the connection when main() exits

	// longURL --> shortURL
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		api.ShortenURLHandler(w, r, redisClient)
	})

	// shortURL --> longURL
	http.HandleFunc("/redirect", api.RedirectHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
