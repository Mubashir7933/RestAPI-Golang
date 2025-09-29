package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Mubashir7933/RestAPI-Golang/internal/types"
	"github.com/Mubashir7933/RestAPI-Golang/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Info("Creating a student")
		response.WriteJSON(w, http.StatusCreated, map[string]string{
			"message": "student created successfully",
		})
	}
}
