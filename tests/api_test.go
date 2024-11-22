package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SartajBhuvaji/api"
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

	// call the handler
	http.HandlerFunc(api.ShortenURLHandler).ServeHTTP(rec, req)

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
