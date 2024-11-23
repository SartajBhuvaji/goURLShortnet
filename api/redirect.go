package api

import (
	"encoding/json"
	"net/http"

	"github.com/SartajBhuvaji/database"
)

type RedirectURLRequest struct {
	LongURL string `json:"long_url"`
}

// handle the short --> long URL redirection

func RedirectHandler(w http.ResponseWriter, r *http.Request, redisClient *database.RedisClient) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the short URL from query parameters
	shortUrl := r.URL.Query().Get("url")
	if shortUrl == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	longUrl, err := redisClient.Get(shortUrl)
	if err != nil {
		http.Error(w, "Error fetching URL from Redis", http.StatusInternalServerError)
		return
	}

	response := RedirectURLRequest{
		LongURL: longUrl,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
