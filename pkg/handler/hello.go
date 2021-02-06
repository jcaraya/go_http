package handler

import (
	"fmt"
	"net/http"
)

// helloPattern defines the URL pattern.
const helloPattern = "/hello"

// HelloHandler provides the handle for hello requests.
type HelloHandler struct{}

// Pattern returns the string pattern of the hello handler.
func (h *HelloHandler) Pattern() string {
	return helloPattern
}

// ServeHTTP handels the http request.
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
