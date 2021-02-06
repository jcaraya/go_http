package main

import (
	"fmt"
	"go_websocket/pkg/handler"
	"go_websocket/pkg/memory"
	"go_websocket/pkg/router"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	fmt.Println("HTTP server started")

	// Initialize a new mux router.
	muxRouter := router.NewMuxRouter()

	// Debug component.
	muxRouter.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	// Register the handlers.
	muxRouter.RegisterHandlers(handler.RegisteredHandlers)

	// Initialize memory consuming component.
	go memory.AllocateMemory(20)

	// Listen for requests.
	http.ListenAndServe(":8090", muxRouter)
}
