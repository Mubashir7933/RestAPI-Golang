package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Mubashir7933/RestAPI-Golang/internal/storage"
	"github.com/Mubashir7933/RestAPI-Golang/internal/types"
	"github.com/Mubashir7933/RestAPI-Golang/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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
		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// log the student data
		slog.Info("user created", slog.String("userId", fmt.Sprint(lastId)))

		response.WriteJSON(w, http.StatusCreated, map[string]int64{
			"id": lastId,
		})
	}
}
