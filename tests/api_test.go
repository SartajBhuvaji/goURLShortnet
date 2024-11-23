package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SartajBhuvaji/api"
	"github.com/SartajBhuvaji/utils"
)

func TestShortenURLHandler(t *testing.T) {
	// create a request body
	reqBody := map[string]string{"url": "https://www.sartaj.me"}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("could not marshal JSON: %v", err)
	}

	// create a request
	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// create a response recorder
	rec := httptest.NewRecorder()

	redisClient, err := utils.SetupRedis()
	if err != nil {
		t.Fatalf("could not set up Redis: %v", err)
	}
	defer redisClient.Close()

	// create a handler function that matches http.HandlerFunc signature
	handler := func(w http.ResponseWriter, r *http.Request) {
		api.ShortenURLHandler(w, r, redisClient)
	}

	http.HandlerFunc(handler).ServeHTTP(rec, req)

	// check the status code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}

	// check the response body
	var respBody map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		t.Fatalf("could not decode JSON: %v", err)
	}

	if _, ok := respBody["short_url"]; !ok {
		t.Fatalf("response body does not have short_url field")
	}
}

func TestRedirectHandler(t *testing.T) {
	// First, create a shortened URL
	redisClient, err := utils.SetupRedis()
	if err != nil {
		t.Fatalf("could not set up Redis: %v", err)
	}
	defer redisClient.Close()

	// Store a test URL in Redis
	testShortURL := "testcode"
	testLongURL := "https://www.example.com"
	err = redisClient.Set(testShortURL, testLongURL)
	if err != nil {
		t.Fatalf("could not set test URL in Redis: %v", err)
	}

	// Create a request to redirect
	req, err := http.NewRequest(http.MethodGet, "/redirect?url="+testShortURL, nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		api.RedirectHandler(w, r, redisClient)
	}

	http.HandlerFunc(handler).ServeHTTP(rec, req)

	// Check status code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}

	// Check response body
	var respBody api.RedirectURLRequest
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		t.Fatalf("could not decode JSON: %v", err)
	}

	if respBody.LongURL != testLongURL {
		t.Errorf("expected long URL %s; got %s", testLongURL, respBody.LongURL)
	}
}

func TestRedirectHandlerInvalidMethod(t *testing.T) {
	redisClient, err := utils.SetupRedis()
	if err != nil {
		t.Fatalf("could not set up Redis: %v", err)
	}
	defer redisClient.Close()

	// Try with POST method instead of GET
	req, err := http.NewRequest(http.MethodPost, "/redirect?url=test", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		api.RedirectHandler(w, r, redisClient)
	}

	http.HandlerFunc(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status Method Not Allowed; got %v", rec.Code)
	}
}

func TestRedirectHandlerMissingURL(t *testing.T) {
	redisClient, err := utils.SetupRedis()
	if err != nil {
		t.Fatalf("could not set up Redis: %v", err)
	}
	defer redisClient.Close()

	// Create request without URL parameter
	req, err := http.NewRequest(http.MethodGet, "/redirect", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		api.RedirectHandler(w, r, redisClient)
	}

	http.HandlerFunc(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status Bad Request; got %v", rec.Code)
	}
}

func TestShortenURLHandlerInvalidJSON(t *testing.T) {
	// Create invalid JSON request body
	invalidJSON := []byte(`{"url": invalid}`)

	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(invalidJSON))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	redisClient, err := utils.SetupRedis()
	if err != nil {
		t.Fatalf("could not set up Redis: %v", err)
	}
	defer redisClient.Close()

	handler := func(w http.ResponseWriter, r *http.Request) {
		api.ShortenURLHandler(w, r, redisClient)
	}

	http.HandlerFunc(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status Bad Request; got %v", rec.Code)
	}
}
