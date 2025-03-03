package models

type Task struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	AssignedUserID string `json:"assigned_user_id"`
}
