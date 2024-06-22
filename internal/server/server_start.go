package server

import (
	"fmt"
	"net/http"

	"github.com/Ra1nz0r/simple_todo_list_restapi/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func StartServe() {
	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetAllTasks)
	r.Post("/tasks", handlers.CreateTask)
	r.Get("/tasks/{id}", handlers.GetTaskById)
	r.Delete("/tasks/{id}", handlers.DeleteTaskById)

	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
