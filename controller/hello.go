package controller

import (
	"net/http"

	"github.com/holzhey/goweb/response"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	res := &response.Hello{
		Message: "Hello world!",
	}
	response.RenderJsonResponse(w, res)
}
