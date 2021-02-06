package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPHandler TODO.
type HTTPHandler interface {
	Pattern() string
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// MuxRouter TODO.
type MuxRouter struct {
	*mux.Router
}

// NewMuxRouter TODO.
func NewMuxRouter() *MuxRouter {
	return &MuxRouter{mux.NewRouter()}
}

// RegisterHandlers TODO.
func (m *MuxRouter) RegisterHandlers(handlers []HTTPHandler) {
	for _, h := range handlers {
		m.Router.Handle(h.Pattern(), h)
	}
}
