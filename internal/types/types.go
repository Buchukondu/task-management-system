package types

type TaskStatus string

const (
	Pending    TaskStatus = "Pending"
	InProgress TaskStatus = "In Progress"
	Completed  TaskStatus = "Completed"
)

type UserRole string

const (
	Admin UserRole = "Admin"
	User  UserRole = "User"
)
