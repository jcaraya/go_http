package handler

import (
	"fmt"
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
	fmt.Fprintf(w, "Default Handler\n")
}
