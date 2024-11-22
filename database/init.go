package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	//"github.com/SartajBhuvaji/database"
	"github.com/joho/godotenv"
)

func databaseInitialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var redisHost string = os.Getenv("REDIS_HOST")
	var redisPassword string = os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error converting REDIS_DB to int: %v", err)
	}

	// Create a new RedisClient
	redisClient := NewRedisClient(redisHost, redisPassword, redisDB)
	//redisClient := database.NewRedisClient(redisHost, redisPassword, redisDB)

	defer redisClient.Close()

	// Test connection
	if err := redisClient.Ping(); err != nil {
		fmt.Printf("Could not connect to Redis: %v\n", err)
		return
	} else {
		fmt.Println("Connected to Redis")
	}

	// Set the counter to 1
	redisClient.SetCounter("counter", 1)
	fmt.Println("Counter set to 1")
}

func main() {
	databaseInitialize()
}

// TO run : go run init.go
