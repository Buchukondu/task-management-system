package main

import (
	"net/http"
	"task-management-system/internal/handlers"
	"task-management-system/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	taskService := services.NewTaskService()
	userService := services.NewUserService()

	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}/assign", taskHandler.AssignTask).Methods("PATCH")

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}/tasks", userHandler.GetUserTasks).Methods("GET")
	r.HandleFunc("/users/{userId}/tasks/{taskId}", userHandler.DeleteUserTask).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
