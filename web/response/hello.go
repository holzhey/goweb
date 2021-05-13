package response

import "net/http"

type Hello struct {
	Message string `json:"message"`
}

func (j *Hello) GetStatus() int {
	return http.StatusOK
}

func (j *Hello) GetPayload() interface{} {
	return j
}
