package main

import (
	"log"

	"github.com/holzhey/goweb/web"
	"github.com/holzhey/goweb/web/controller"
)

func main() {
	web.AddRoute("/", controller.HandleHello)
	log.Fatal(web.ListenAndServe(":8080"))
}
