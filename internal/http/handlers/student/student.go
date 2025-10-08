package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

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
func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("GetById handler called", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, e := storage.GetStudentById(intId)
		if e != nil {
			slog.Error("Error while fetching the student id", slog.String("id", id), slog.String("error", e.Error()))
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(e))
			return
		}
		response.WriteJSON(w, http.StatusOK, student)

	}
}
func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, errs := storage.GetStudents()
		if errs != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(errs))
			return
		}
		response.WriteJSON(w, http.StatusOK, students)
	}
}
