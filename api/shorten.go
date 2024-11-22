package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/SartajBhuvaji/database"
	"github.com/SartajBhuvaji/utils"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type ShortenURLRequest struct {
	URL string `json:"url"`
}

type ShortenURLResponse struct {
	ShortURL string `json:"short_url"`
}

// handle the URL shortening request
func ShortenURLHandler(w http.ResponseWriter, r *http.Request, redisClient *database.RedisClient) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	counter, err := redisClient.GetCounter() // flag
	if err != nil {
		http.Error(w, "Error fetching counter from Reddis", http.StatusInternalServerError)
		return
	}

	log.Println("Counter: ", counter)

	enc := EncodeURL(counter)
	shortURL := "something/" + enc

	// Update the counter counter++
	err = redisClient.SetCounter("counter", counter+1)

	// Add the short URL to the database
	err = redisClient.Set(shortURL, req.URL)

	if err != nil {
		http.Error(w, "Error updating counter in Reddis", http.StatusInternalServerError)
		return
	}

	resp := ShortenURLResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func EncodeURL(no int) string {
	if no == 0 {
		return string(base62Chars[0])
	}

	base := len(base62Chars)
	var encoded strings.Builder

	for no > 0 {
		remainder := no % base
		encoded.WriteByte(base62Chars[remainder])
		no /= base
	}

	return utils.ReverseString(encoded.String())
}
