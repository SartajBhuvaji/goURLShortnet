package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/SartajBhuvaji/database"
	"github.com/joho/godotenv"
)

// ReverseString reverses a given string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SetupRedis() (*database.RedisClient, error) {
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
	redisClient := database.NewRedisClient(redisHost, redisPassword, redisDB)

	if err := redisClient.Ping(); err != nil {
		fmt.Printf("Could not connect to Redis: %v\n", err)
		return nil, err
	} else {
		fmt.Println("Connected to Redis")
		return redisClient, nil
	}
}
