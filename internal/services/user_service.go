package services

import (
	"errors"
	"strconv"
	"sync"
	"task-management-system/internal/models"
)

type UserService struct {
	users []models.User
	mu    sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		users: []models.User{},
	}
}

func (us *UserService) CreateUser(user models.User) models.User {
	us.mu.Lock()
	defer us.mu.Unlock()
	user.ID = strconv.Itoa(len(us.users) + 1)
	us.users = append(us.users, user)
	return user
}

func (us *UserService) GetUser(id string) (models.User, error) {
	us.mu.Lock()
	defer us.mu.Unlock()
	for _, user := range us.users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (us *UserService) GetUserTasks(userID string) []models.Task {
	us.mu.Lock()
	defer us.mu.Unlock()
	for _, user := range us.users {
		if user.ID == userID {
			return user.Tasks
		}
	}
	return []models.Task{}
}

func (us *UserService) DeleteUserTask(userID string, taskID string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()
	for i, user := range us.users {
		if user.ID == userID {
			for j, task := range user.Tasks {
				if task.ID == taskID {
					us.users[i].Tasks = append(us.users[i].Tasks[:j], us.users[i].Tasks[j+1:]...)
					return true
				}
			}
		}
	}
	return false
}
