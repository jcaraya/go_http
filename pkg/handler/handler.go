package handler

import "go_websocket/pkg/router"

var RegisteredHandlers = []router.HTTPHandler{
	&HelloHandler{},
	&HeaderHandler{},
	&DefaultHandler{},
}
