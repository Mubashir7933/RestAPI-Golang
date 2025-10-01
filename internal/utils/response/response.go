package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`          // showing in postman
	Error  string `json:"error,omitempty"` //omitempty means if there is no error it won't show anything
}

const (
	StatusOK    = "OK"
	StatusError = "error"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}
