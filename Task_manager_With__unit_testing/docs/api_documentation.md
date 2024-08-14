

## TaskController Documentation

### Overview

The `TaskController` is responsible for handling HTTP requests related to task management and user operations in the Task Manager API. It ensures that tasks are managed correctly and that user-related operations are properly handled.

### Endpoints

#### Task Operations

1. **Get All Tasks**

   - **Endpoint:** `GET /tasks`
   - **Description:** Retrieves a list of all tasks for the authenticated user.
   - **Response Status Code:** `200 OK`
   - **Response Body Example:**
     ```json
     [
       {
         "id": "1",
         "title": "Task 1",
         "description": "Description of Task 1",
         "due_date": "2024-08-15T00:00:00Z",
         "status": "pending",
         "owner_id": "user1"
       },
       {
         "id": "2",
         "title": "Task 2",
         "description": "Description of Task 2",
         "due_date": "2024-08-16T00:00:00Z",
         "status": "completed",
         "owner_id": "user1"
       }
     ]
     ```

2. **Get Task By ID**

   - **Endpoint:** `GET /tasks/:id`
   - **Description:** Retrieves a specific task by its ID for the authenticated user.
   - **Response Status Code:** `200 OK`
   - **Response Body Example:**
     ```json
     {
       "id": "1",
       "title": "Task 1",
       "description": "Description of Task 1",
       "due_date": "2024-08-15T00:00:00Z",
       "status": "pending",
       "owner_id": "user1"
     }
     ```

3. **Add Task**

   - **Endpoint:** `POST /tasks`
   - **Description:** Creates a new task for the authenticated user.
   - **Request Body Example:**
     ```json
     {
       "title": "New Task",
       "description": "Description of the new task",
       "due_date": "2024-08-20T00:00:00Z",
       "status": "pending"
     }
     ```
   - **Response Status Code:** `201 Created`
   - **Response Body Example:**
     ```json
     {
       "id": "3",
       "title": "New Task",
       "description": "Description of the new task",
       "due_date": "2024-08-20T00:00:00Z",
       "status": "pending",
       "owner_id": "user1"
     }
     ```

4. **Update Task**

   - **Endpoint:** `PUT /tasks/:id`
   - **Description:** Updates an existing task by its ID for the authenticated user.
   - **Request Body Example:**
     ```json
     {
       "title": "Updated Task Title",
       "description": "Updated description",
       "due_date": "2024-08-22T00:00:00Z",
       "status": "in-progress"
     }
     ```
   - **Response Status Code:** `200 OK`
   - **Response Body Example:**
     ```json
     {
       "id": "1",
       "title": "Updated Task Title",
       "description": "Updated description",
       "due_date": "2024-08-22T00:00:00Z",
       "status": "in-progress",
       "owner_id": "user1"
     }
     ```

5. **Delete Task**

   - **Endpoint:** `DELETE /tasks/:id`
   - **Description:** Deletes a task by its ID for the authenticated user.
   - **Response Status Code:** `200 OK`

#### User Operations

1. **Register User**

   - **Endpoint:** `POST /register`
   - **Description:** Registers a new user.
   - **Request Body Example:**
     ```json
     {
       "email": "testuser@example.com",
       "password": "password"
     }
     ```
   - **Response Status Code:** `200 OK`
   - **Response Body Example:**
     ```json
     {
       "message": "User registered successfully"
     }
     ```

2. **Login User**

   - **Endpoint:** `POST /login`
   - **Description:** Authenticates a user and returns a token.
   - **Request Body Example:**
     ```json
     {
       "email": "testuser@example.com",
       "password": "password"
     }
     ```
   - **Response Status Code:** `200 OK`
   - **Response Body Example:**
     ```json
     {
       "token": "your_jwt_token_here"
     }
     ```



### Test Cases

#### Task Operations

1. **TestTaskController_GetTasks**

   - **Description:** Verifies that the `GetTasks` method returns all tasks for the authenticated user.
   - **Expected Outcome:** Response status `200 OK` and a JSON array of tasks.

2. **TestTaskController_GetTasksById**

   - **Description:** Verifies that the `GetTasksById` method returns a specific task by ID for the authenticated user.
   - **Expected Outcome:** Response status `200 OK` and a JSON object of the task.

3. **TestTaskController_AddTask**

   - **Description:** Verifies that the `AddTask` method creates a new task successfully for the authenticated user.
   - **Expected Outcome:** Response status `201 Created` and a JSON object of the created task.

4. **TestTaskController_UpdateTask**

   - **Description:** Verifies that the `UpdateTask` method updates an existing task successfully for the authenticated user.
   - **Expected Outcome:** Response status `200 OK` and a JSON object of the updated task.

5. **TestTaskController_DeleteTask**

   - **Description:** Verifies that the `DeleteTask` method deletes a task successfully for the authenticated user.
   - **Expected Outcome:** Response status `200 OK`

#### User Operations

1. **TestUserController_Register**

   - **Description:** Verifies that the `Register` method creates a new user successfully.
   - **Expected Outcome:** Response status `200 OK` with a confirmation message.

2. **TestUserController_Login**

   - **Description:** Verifies that the `Login` method authenticates a user and returns a token.
   - **Expected Outcome:** Response status `200 OK` with a JWT token.


### User-Based Task Management

The `TaskController` ensures that tasks and user operations are managed on a per-user basis:

- Users can only view, update, or delete tasks they own.
- Admins have additional permissions to manage users and view all tasks.

### Conclusion

This documentation outlines the functionalities of the `TaskController`, including task and user operations. It provides detailed descriptions of each endpoint, examples of requests and responses, and test cases to verify the correctness of the controller’s operations.

