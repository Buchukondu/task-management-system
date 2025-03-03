package handlers

import (
	"encoding/json"
	"net/http"
	"task-management-system/internal/models"
	"task-management-system/internal/services"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdUser := h.service.CreateUser(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	tasks := h.service.GetUserTasks(userID)
	json.NewEncoder(w).Encode(tasks)
}

func (h *UserHandler) DeleteUserTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	taskID := vars["taskId"]
	if h.service.DeleteUserTask(userID, taskID) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "task not found", http.StatusNotFound)
	}
}
