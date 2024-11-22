package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient is a struct that encapsulates the Redis client.
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

type URLValue struct {
	OriginalURL  string `json:"url"`
	CreatedAt    string `json:"created_at"`
	LastAccessed string `json:"last_accessed"`
	AccessCount  int    `json:"access_count"`
}

// Implement the encoding.BinaryMarshaler interface
func (u URLValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// Implement the encoding.BinaryUnmarshaler interface
func (u *URLValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

// NewRedisClient initializes a new Redis client.
func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db, // Use default DB
	})
	return &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Ping checks the connection to the Redis server.
func (r *RedisClient) Ping() error {
	pong, err := r.client.Ping(r.ctx).Result()
	if err != nil {
		return err
	}
	fmt.Printf("Connected to Redis: %s\n", pong)
	return nil
}

// Set counter
func (r *RedisClient) SetCounter(key string, value int) error {
	err := r.client.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	fmt.Printf("Key '%s' set successfully!\n", key)
	return nil
}

// Get Counter
func (r *RedisClient) GetCounter() (int, error) {
	val, err := r.client.Get(r.ctx, "counter").Result()
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, err
	}
	return intVal, nil
}

// Set stores a key-value pair in Redis.
func (r *RedisClient) Set(originalURL, shortURL string) error {

	value := URLValue{
		OriginalURL:  originalURL,
		CreatedAt:    time.Now().String(),
		LastAccessed: time.Now().String(),
		AccessCount:  1,
	}

	err := r.client.Set(r.ctx, shortURL, value, 0).Err()
	if err != nil {
		fmt.Printf("Unable to add new key '%s' to Redis: %v\n", shortURL, err)
		return err
	}
	fmt.Printf("Key '%s' set successfully!\n", shortURL)
	return nil
}

// Get retrieves the original URL of a given key from Redis.
func (r *RedisClient) Get(shortURL string) (string, error) {

	var urlValue URLValue
	err := r.client.Get(r.ctx, shortURL).Scan(&urlValue)
	if err != nil {
		log.Println("Error fetching URL from Redis: ", err)
		return "", err
	}

	urlValue.LastAccessed = time.Now().String()
	urlValue.AccessCount++

	r.client.Set(r.ctx, shortURL, urlValue, 0).Err()
	return urlValue.OriginalURL, nil
}

// Close closes the Redis client connection.
func (r *RedisClient) Close() {
	err := r.client.Close()
	if err != nil {
		fmt.Printf("Error closing Redis connection: %v\n", err)
	} else {
		fmt.Println("Redis connection closed.")
	}
}
