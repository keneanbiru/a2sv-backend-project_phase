
## API Documentation

### **1. User Management**

#### **Register a New User**

- **Endpoint:** `POST /register`
- **Description:** Registers a new user in the system.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string",
    "role": "string"  // Optional: can be "admin" or "user"
  }
  ```
- **Response:**
  - **Status Code:** `201 Created`
  - **Body:**
    ```json
    {
      "id": "object_id",
      "username": "string",
      "role": "string"
    }
    ```
- **Roles:** No role required for registration if role specified its for admin.
- **Notes:** Default role is "user" if not specified.

#### **Login**

- **Endpoint:** `POST /login`
- **Description:** Authenticates a user and returns a JWT token.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    {
      "token": "jwt_token"
    }
    ```
- **Roles:** No role required for login.
- **Notes:** The JWT token must be included in the Authorization header for protected routes.

---

### **2. Authentication and Authorization Middleware**

- **Purpose:** Ensures that requests to protected routes are authenticated and authorized.
- **Logic:**
  - **Token Extraction:** Extracts the JWT token from the Authorization header.
  - **Token Verification:** Validates the token using the secret key.
  - **Context Population:** Adds the user's ID and role to the request context (`c.Set("ID", userID)` and `c.Set("Role", userRole)`), which can be accessed in subsequent handlers.

---

### **3. Task Management**

#### **Create a Task**

- **Endpoint:** `POST /tasks`
- **Description:** Creates a new task.
- **Request Body:**
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "yyyy-mm-ddTHH:MM:SSZ",
    "status": "string" // Optional
  }
  ```
- **Response:**
  - **Status Code:** `201 Created`
  - **Body:**
    ```json
    {
      "id": "object_id",
      "title": "string",
      "description": "string",
      "due_date": "yyyy-mm-ddTHH:MM:SSZ",
      "status": "string",
      "owner_id": "object_id"
    }
    ```
- **Roles:**
  - **Admin:** Can assign tasks to any user by setting the `owner_id`.
  - **Normal User:** Can only create tasks for themselves (the `owner_id` is set to the current user's ID).
- **Notes:** Admins can set the `owner_id` to assign tasks to other users. Normal users cannot specify `owner_id` and thus can only create tasks for themselves.

#### **Retrieve Tasks**

- **Endpoint:** `GET /tasks`
- **Description:** Retrieves tasks based on the user’s role and ID.
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    [
      {
        "id": "object_id",
        "title": "string",
        "description": "string",
        "due_date": "yyyy-mm-ddTHH:MM:SSZ",
        "status": "string",
        "owner_id": "object_id"
      }
    ]
    ```
- **Roles:**
  - **Admin:** Can retrieve all tasks.
  - **Normal User:** Can retrieve only their own tasks.
- **Notes:** Admins can see all tasks while normal users only see their own tasks.

#### **Update a Task**

- **Endpoint:** `PUT /tasks/:task_id`
- **Description:** Updates a specific task.
- **Request Body:**
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "yyyy-mm-ddTHH:MM:SSZ",
    "status": "string" // Optional
  }
  ```
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    {
      "modified_count": "number"
    }
    ```
- **Roles:**
  - **Admin:** Can update any task.
  - **Normal User:** Can only update tasks they own.
- **Notes:** Admins have full access to update any task, while normal users can only update their own tasks.

#### **Delete a Task**

- **Endpoint:** `DELETE /tasks/:task_id`
- **Description:** Deletes a specific task.
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    {
      "deleted_count": "number"
    }
    ```
- **Roles:**
  - **Admin:** Can delete any task.
  - **Normal User:** Can only delete tasks they own.
- **Notes:** Admins can delete any task, while normal users can only delete their own tasks.

---

### **Security Considerations**

- **Password Storage:** Ensure passwords are hashed before storing them in the database.
- **Token Security:** Use strong secret keys and secure token handling practices to prevent unauthorized access.
- **Role-Based Access Control:** Ensure that routes and operations are properly restricted based on user roles.


