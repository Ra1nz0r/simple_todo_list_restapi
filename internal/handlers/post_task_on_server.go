package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ra1nz0r/simple_todo_list_restapi/internal/globals"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task globals.Task

	result, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(result, &task); err != nil {
		errorMsg := fmt.Sprintf("Can't deserialize: %s", err.Error())
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	_, ok := globals.Tasks[task.ID]
	if ok {
		errorMsg := fmt.Sprintf("Task #%s already exists.", task.ID)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	globals.Tasks[task.ID] = task

	w.WriteHeader(http.StatusCreated)
}
