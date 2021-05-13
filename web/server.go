package web

import (
	"fmt"
	"net/http"
)

type route struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
}

var routes []route

func AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	for _, route := range routes {
		if route.pattern == pattern {
			panic(fmt.Sprintf("pattern %s already added", pattern))
		}
	}
	routes = append(routes, route{pattern, handler})
}

func ListenAndServe(port string) error {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.pattern, route.handler)
	}

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	return server.ListenAndServe()
}
