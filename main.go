package main

import (
	"log"
	"net/http"

	"github.com/holzhey/goweb/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.HandleHello)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
