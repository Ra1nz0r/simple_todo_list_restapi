package handlers

import (
	"net/http"

	"github.com/Ra1nz0r/simple_todo_list_restapi/internal/globals"

	"github.com/go-chi/chi/v5"
)

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskStruct, ok := globals.Tasks[id]
	if !ok {
		http.Error(w, "Task not found.", http.StatusBadRequest)
		return
	}

	delete(globals.Tasks, taskStruct.ID)

	w.WriteHeader(http.StatusOK)
}
