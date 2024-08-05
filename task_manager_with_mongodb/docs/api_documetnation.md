# Task Manager API Documentation

## Overview

The Task Manager API is a simple RESTful service implemented using Go and the Gin Framework. It provides endpoints for managing tasks, including creating, reading, updating, and deleting tasks. The tasks are stored in an in-memory database, and proper error handling is implemented for various scenarios.

## Requirements

- **Language**: Go
- **Framework**: Gin
- **Database**: In-memory storage

### Using Postman

You can find a detailed Postman collection documenting all the API endpoints, including request and response examples, [here](https://documenter.getpostman.com/view/37367045/2sA3rxpYeB)


1. **Set up Postman**: Open Postman and set the base URL for your requests (e.g., `http://localhost:8080`).
2. **Create Requests**: Set up requests for each endpoint (GET, POST, PUT, DELETE) with the appropriate URL, method, headers, and body.
3. **Authorization**: If your API requires authorization, set the necessary headers or tokens.
4. **Send Requests**: Send requests and verify the responses.

## Endpoints

### 1. GET /tasks

**Description**: Retrieve a list of all tasks.

- **Method**: GET
- **URL**: `/tasks`
- **Request Parameters**: None
- **Response**:
  - **200 OK**: Successfully retrieved the list of tasks.
  - **Response Body**:
    ```json
    [
      {
        "id": "task-1",
        "title": "Sample Task",
        "description": "This is a sample task",
        "due_date": "2024-08-01T00:00:00Z",
        "status": "Pending"
      },
      ...
    ]
    ```

### 2. GET /tasks/:id

**Description**: Retrieve the details of a specific task by ID.

- **Method**: GET
- **URL**: `/tasks/:id`
- **Request Parameters**:
  - **Path**: `:id` - The ID of the task
- **Response**:
  - **200 OK**: Successfully retrieved the task details.
  - **404 Not Found**: Task with the specified ID not found.
  - **Response Body (200 OK)**:
    ```json
    {
      "id": "task-1",
      "title": "Sample Task",
      "description": "This is a sample task",
      "due_date": "2024-08-01T00:00:00Z",
      "status": "Pending"
    }
    ```

### 3. POST /tasks

**Description**: Create a new task.

- **Method**: POST
- **URL**: `/tasks`
- **Request Body**:
  - **title**: `string` (required)
  - **description**: `string` (optional)
  - **due_date**: `string` (required, in ISO 8601 format)
  - **status**: `string` (required)
- **Response**:
  - **201 Created**: Successfully created a new task.
  - **400 Bad Request**: Invalid input data.
  - **Response Body (201 Created)**:
    ```json
    {
      "id": "task-2",
      "title": "New Task",
      "description": "Description of the new task",
      "due_date": "2024-08-15T00:00:00Z",
      "status": "Pending"
    }
    ```

### 4. PUT /tasks/:id

**Description**: Update a specific task.

- **Method**: PUT
- **URL**: `/tasks/:id`
- **Request Parameters**:
  - **Path**: `:id` - The ID of the task
- **Request Body**:
  - **title**: `string` (optional)
  - **description**: `string` (optional)
  - **due_date**: `string` (optional, in ISO 8601 format)
  - **status**: `string` (optional)
- **Response**:
  - **200 OK**: Successfully updated the task.
  - **404 Not Found**: Task with the specified ID not found.
  - **400 Bad Request**: Invalid input data.
  - **Response Body (200 OK)**:
    ```json
    {
      "id": "task-1",
      "title": "Updated Task",
      "description": "Updated description",
      "due_date": "2024-08-15T00:00:00Z",
      "status": "Completed"
    }
    ```

### 5. DELETE /tasks/:id

**Description**: Delete a specific task.

- **Method**: DELETE
- **URL**: `/tasks/:id`
- **Request Parameters**:
  - **Path**: `:id` - The ID of the task
- **Response**:
  - **204 No Content**: Successfully deleted the task.
  - **404 Not Found**: Task with the specified ID not found.

## Error Handling

- **400 Bad Request**: Returned when the request body contains invalid data or is missing required fields.
- **404 Not Found**: Returned when a task with the specified ID is not found.
- **500 Internal Server Error**: Returned in case of an unexpected error on the server.

## Folder Structure

```
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

- **main.go**: Entry point of the application. Initializes the router and starts the server.
- **controllers/task_controller.go**: Contains handler functions for HTTP requests.
- **models/task.go**: Defines the `Task` struct representing the data model.
- **data/task_service.go**: Contains business logic and data manipulation methods for tasks.
- **router/router.go**: Configures and sets up the routing for the API using Gin.
- **docs/api_documentation.md**: Contains API documentation.

## Testing the API

### Sample Data for Testing

**Creating a Task**:

- **POST /tasks**
  - **Request Body**:
    ```json
    {
      "title": "Learn Go",
      "description": "Complete the Go course on Coursera",
      "due_date": "2024-08-15T00:00:00Z",
      "status": "In Progress"
    }
    ```

**Updating a Task**:

- **PUT /tasks/task-1**
  - **Request Body**:
    ```json
    {
      "title": "Learn Go",
      "description": "Complete the Go course and practice exercises",
      "due_date": "2024-08-20T00:00:00Z",
      "status": "In Progress"
    }
    ```

## Instructions for Running the API

1. **Clone the Repository**: Clone the project repository to your local machine.
2. **Install Dependencies**: Navigate to the project directory and run `go mod tidy` to install the necessary dependencies.
3. **Run the Server**: Use `go run main.go` to start the server. The server will listen on `http://localhost:8080` by default.

## Notes

- The current implementation uses an in-memory database, meaning all data will be lost when the server is stopped. Future updates may include persistent storage.
- Ensure proper error handling and validation for all API endpoints.
- Follow Go best practices for clean, maintainable, and efficient code.

This documentation provides an overview of the API's functionality, endpoints, and usage. For further details and updates, refer to the `docs/api_documentation.md` file and the source code.
