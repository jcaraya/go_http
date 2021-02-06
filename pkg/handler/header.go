package handler

import (
	"fmt"
	"net/http"
)

// headerPattern defines the URL pattern.
const headerPattern = "/header"

// HeaderHandler provides the handle for header requests.
type HeaderHandler struct{}

// Pattern returns the string pattern of the header handler.
func (h *HeaderHandler) Pattern() string {
	return headerPattern
}

// ServeHTTP handels the http request.
func (h *HeaderHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
