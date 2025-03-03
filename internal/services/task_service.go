package services

import (
	"errors"
	"strconv"
	"sync"
	"task-management-system/internal/models"
)

type TaskService struct {
	tasks []models.Task
	mu    sync.Mutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []models.Task{},
	}
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.ID = strconv.Itoa(len(s.tasks) + 1)
	s.tasks = append(s.tasks, task)
	return task
}

func (s *TaskService) GetTask(id string) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (s *TaskService) UpdateTask(updatedTask models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, task := range s.tasks {
		if task.ID == updatedTask.ID {
			s.tasks[i] = updatedTask
			return updatedTask, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (s *TaskService) DeleteTask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) AssignTask(taskID string, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, task := range s.tasks {
		if task.ID == taskID {
			s.tasks[i].AssignedUserID = userID
			return nil
		}
	}
	return errors.New("task not found")
}
