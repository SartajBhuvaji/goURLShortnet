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
	shortUrl := req.URL
	longUrl, err := redisClient.Get(shortUrl)
	if err != nil {
		http.Error(w, "Error fetching URL from Reddis", http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, longUrl, http.StatusMovedPermanently)

	response := RedirectURLRequest{
		LongURL: longUrl,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
