package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SartajBhuvaji/api"
	"github.com/SartajBhuvaji/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var redisHost string
	var redisPassword string

	// // Read from environment variables
	redisHost = os.Getenv("REDIS_HOST")
	redisPassword = os.Getenv("REDIS_PASSWORD")

	// Create a new RedisClient

	redisClient := database.NewRedisClient(redisHost, redisPassword, 0)
	defer redisClient.Close()

	// Test connection
	if err := redisClient.Ping(); err != nil {
		fmt.Printf("Could not connect to Redis: %v\n", err)
		return
	} else {
		fmt.Println("Connected to Redis")
	}

	// longURL --> shortURL
	http.HandleFunc("/shorten", api.ShortenURLHandler)

	// shortURL --> longURL
	http.HandleFunc("/redirect", api.RedirectHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
