# Task Management System

This is a simple task management system implemented in Go. It allows you to create, get, update, delete, and assign tasks to users.

## Endpoints

- `POST /tasks` - Create a new task
- `GET /tasks/{id}` - Get a task by ID
- `PUT /tasks/{id}` - Update a task by ID
- `DELETE /tasks/{id}` - Delete a task by ID
- `PATCH /tasks/{id}/assign` - Assign a task to a user
- `GET /users/{id}/tasks` - Get tasks assigned to a user
- `DELETE /users/{userId}/tasks/{taskId}` - Delete a task assigned to a user

## Project Structure

```
task-management-system
├── cmd
│   └── main.go               # Entry point of the application
├── internal
│   ├── handlers
│   │   ├── task_handler.go    # Handles HTTP requests for tasks
│   │   └── user_handler.go     # Handles HTTP requests for users
│   ├── models
│   │   ├── task.go            # Defines the Task struct
│   │   └── user.go            # Defines the User struct
│   ├── services
│   │   ├── task_service.go     # Manages task operations
│   │   └── user_service.go      # Manages user operations
│   └── types
│       └── types.go           # Common types and constants
├── go.mod                     # Module definition
└── README.md                  # Project documentation
```

## Running the Application

1. Install dependencies:
    ```sh
    go mod tidy
    ```

2. Run the application:
    ```sh
    go run cmd/main.go
    ```

The application will be running on [http://localhost:8080](http://_vscodecontentref_/10).

## License

This project is licensed under the MIT License.