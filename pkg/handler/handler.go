package handler

import "go_http/pkg/router"

var RegisteredHandlers = []router.HTTPHandler{
	&HelloHandler{},
	&HeaderHandler{},
	&DefaultHandler{},
}
