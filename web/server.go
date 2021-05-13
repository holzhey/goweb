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

func AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) error {
	for _, route := range routes {
		if route.pattern == pattern {
			return fmt.Errorf("pattern %s already added as a route", pattern)
		}
	}
	routes = append(routes, route{pattern, handler})
	return nil
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
