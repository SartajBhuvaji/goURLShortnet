package api

import (
	"net/http"
)

// handle the short --> long URL redirection
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
