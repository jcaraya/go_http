package handler

import (
	"encoding/json"
	"net/http"
)

const defaultPattern = "/"

// DefaultHandler provides the handle for default requests.
type DefaultHandler struct{}

// Pattern returns the string pattern of the hello handler.
func (h *DefaultHandler) Pattern() string {
	return defaultPattern
}

// ServeHTTP handels the http request.
func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Initialize an empty slice that will contain the available handlers.
	p := []string{}

	// Iterate over the available handlers.
	for _, h := range RegisteredHandlers {
		p = append(p, h.Pattern())
	}

	// Marshal the resulting list.
	r, err := json.Marshal(p)
	if err != nil {
		// TODO: log error.
		return
	}

	// Write the corresponding json response.
	_, err = w.Write(r)
	if err != nil {
		// TODO: log error.
		return
	}
}
