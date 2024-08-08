package controllers

import (
	"net/http"
	"time"
	"user_authentication/data"
	"user_authentication/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JwtSecret = []byte("your_secret_key")

type TaskController struct {
	taskService *data.TaskService
}

type UserController struct {
	userService *data.UserService
}

func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func NewUserController(userService *data.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register handles user registration
func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user.ID = primitive.NewObjectID()

	if user.Role == "" {
		user.Role = "user" // Default role
	}

	if err := uc.userService.RegisterUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login handles user login and returns a JWT token
func (uc *UserController) Login(c *gin.Context) {
	var user_info struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uc.userService.LoginUser(user_info.Username, user_info.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(100 * time.Minute)
	claims := &models.Claims{
		ID:   user.ID.Hex(), // Convert ObjectID to string
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// AddTask handles task creation
func (tc *TaskController) AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userIDStr, exists := c.Get("ID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := c.GetString("Role")

	// Convert userIDStr from string to primitive.ObjectID
	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	if newTask.OwnerID == primitive.NilObjectID {
		// If OwnerID is not provided, set to the current user's ID
		newTask.OwnerID = userID
	} else {
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can assign tasks to other users"})
			return
		}

		// Verify that the specified OwnerID corresponds to an existing user
		ownerExists, err := tc.taskService.UserExists(newTask.OwnerID)
		if err != nil || !ownerExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Specified user does not exist"})
			return
		}
	}

	// Call the service layer to create the task
	result, err := tc.taskService.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task: " + err.Error()})
		return
	}

	if objid, ok := result.InsertedID.(primitive.ObjectID); ok {
		newTask.ID = objid
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inserted task ID"})
		return
	}
	
	c.JSON(http.StatusCreated, newTask)
}

// GetTasks retrieves tasks based on user ID and role
func (tc *TaskController) GetTasks(c *gin.Context) {
	userIDStr, exists := c.Get("ID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := c.GetString("Role")

	// Convert userIDStr from string to primitive.ObjectID
	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	// Retrieve tasks based on the user's role
	tasks, err := tc.taskService.Gettasks(userID, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// UpdateTask handles task updates
func (tc *TaskController) UpdateTask(c *gin.Context) {
	taskIDStr := c.Param("task_id")
	taskID, err := primitive.ObjectIDFromHex(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID = taskID
	result, err := tc.taskService.UpdateTask(taskID, updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"modified_count": result.ModifiedCount})
}

// DeleteTask handles task deletion
func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskIDStr := c.Param("task_id")
	taskID, err := primitive.ObjectIDFromHex(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	result, err := tc.taskService.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_count": result.DeletedCount})
}



































// package controllers

// import (
// 	"net/http"
// 	"time"
// 	"user_authentication/data"
// 	"user_authentication/models"

// 	// "user_authentication/config"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// var JwtSecret = []byte("your_secret_key")

// type TaskController struct {
// 	taskService *data.TaskService
// }

// type UserController struct {
// 	userService *data.UserService
// }

// func NewTaskController(taskService *data.TaskService) *TaskController {
// 	return &TaskController{taskService: taskService}
// }

// func NewUserController(userService *data.UserService) *UserController {
// 	return &UserController{userService: userService}
// }

// // Register handles user registration
// func (uc *UserController) Register(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}
// 	user.ID = primitive.NewObjectID()

// 	if user.Role == "" {
// 		user.Role = "user" // Default role
// 	}

// 	if err := uc.userService.RegisterUser(user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, user)
// }

// // Login handles user login and returns a JWT token
// func (uc *UserController) Login(c *gin.Context) {
// 	var user_info struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&user_info); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	user, err := uc.userService.LoginUser(user_info.Username, user_info.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	expirationTime := time.Now().Add(100 * time.Minute)
// 	claims := &models.Claims{
// 		ID:   user.ID.Hex(), // Convert ObjectID to string
// 		Role: user.Role,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(JwtSecret)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": tokenString})
// }

// // AddTask handles task creation
// func (tc *TaskController) AddTask(c *gin.Context) {
// 	var newTask models.Task
// 	if err := c.ShouldBindJSON(&newTask); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	userIDStr, exists := c.Get("ID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	role := c.GetString("Role")

// 	// Convert userIDStr from string to primitive.ObjectID
// 	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
// 		return
// 	}

// 	if newTask.OwnerID == primitive.NilObjectID {
// 		// If OwnerID is not provided, set to the current user's ID
// 		newTask.OwnerID = userID
// 	} else {
// 		if role != "admin" {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can assign tasks to other users"})
// 			return
// 		}

// 		// Verify that the specified OwnerID corresponds to an existing user
// 		ownerExists, err := tc.taskService.UserExists(newTask.OwnerID)
// 		if err != nil || !ownerExists {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Specified user does not exist"})
// 			return
// 		}
// 	}

// 	// Call the service layer to create the task
// 	result, err := tc.taskService.CreateTask(newTask)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task: " + err.Error()})
// 		return
// 	}

// 	if objid, ok := result.InsertedID.(primitive.ObjectID); ok {
// 		newTask.ID = objid
// 	} else {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inserted task ID"})
// 		return
// 	}
	
// 	c.JSON(http.StatusCreated, newTask)
// }

// // GetTasks retrieves tasks based on user ID and role
// func (tc *TaskController) GetTasks(c *gin.Context) {
// 	userIDStr, exists := c.Get("ID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	role := c.GetString("Role")

// 	// Convert userIDStr from string to primitive.ObjectID
// 	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
// 		return
// 	}

// 	// Retrieve tasks based on the user's role
// 	tasks, err := tc.taskService.Gettasks(userID, role)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, tasks)
// }
