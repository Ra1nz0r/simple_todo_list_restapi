package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ra1nz0r/simple_todo_list_restapi/internal/globals"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(globals.Tasks)
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
