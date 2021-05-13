package response

import (
	"encoding/json"
	"net/http"
)

type JsonResponse interface {
	GetStatus() int
	GetPayload() interface{}
}

func RenderJsonResponse(w http.ResponseWriter, p JsonResponse) {
	w.WriteHeader(p.GetStatus())
	json.NewEncoder(w).Encode(p.GetPayload())
}
