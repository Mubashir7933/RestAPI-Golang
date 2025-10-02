package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Mubashir7933/RestAPI-Golang/internal/types"
	"github.com/Mubashir7933/RestAPI-Golang/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
		}

		//request validation
		validator.New().Struct(student)

		// log the student data
		slog.Info("Student Data", slog.Any("student", student))

		slog.Info("Creating a student")
		response.WriteJSON(w, http.StatusCreated, map[string]string{
			"message": "student created successfully",
		})
	}
}
