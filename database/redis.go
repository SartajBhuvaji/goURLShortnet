package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisClient is a struct that encapsulates the Redis client.
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
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

// Set stores a key-value pair in Redis.
func (r *RedisClient) Set(key, value string) error {
	err := r.client.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	fmt.Printf("Key '%s' set successfully!\n", key)
	return nil
}

// Get retrieves the value of a given key from Redis.
func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
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

// // Main function
// func main() {

// 	// Test connection
// 	if err := redisClient.Ping(); err != nil {
// 		fmt.Printf("Could not connect to Redis: %v\n", err)
// 		return
// 	}

// 	// Set a key-value pair
// 	if err := redisClient.Set("mykey", "myvalue"); err != nil {
// 		fmt.Printf("Error setting key: %v\n", err)
// 		return
// 	}

// 	// Get the value of a key
// 	value, err := redisClient.Get("mykey")
// 	if err != nil {
// 		fmt.Printf("Error getting key: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Retrieved value: %s\n", value)
// }
