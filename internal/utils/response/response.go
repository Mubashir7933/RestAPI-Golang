package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsges []string

	for _, e := range errs {
		switch e.ActualTag() {
		case "required":
			errMsges = append(errMsges, fmt.Sprintf("field %s is required field", e.Field()))
		default:
			errMsges = append(errMsges, fmt.Sprintf("field %s is not valid", e.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsges, ", "),
	}
}
