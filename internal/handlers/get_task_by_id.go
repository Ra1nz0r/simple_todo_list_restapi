package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ra1nz0r/simple_todo_list_restapi/internal/globals"

	"github.com/go-chi/chi/v5"
)

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskStruct, ok := globals.Tasks[id]
	if !ok {
		http.Error(w, "Task not found.", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(taskStruct)
	if err != nil {
		errorMsg := fmt.Sprintf("Can't serialize: %s", err.Error())
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, "Error.", http.StatusBadRequest)
		return
	}
}
