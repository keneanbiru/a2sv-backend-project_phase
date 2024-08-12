package controllers

import (
	"net/http"
	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskusecase domain.TaskUsecase
	userusecase domain.UserUsecase
}

// NewTaskController creates a new instance of TaskController with the provided task and user usecases.
func NewTaskController(taskmgr domain.TaskUsecase, usermgr domain.UserUsecase) *TaskController {
	return &TaskController{
		taskusecase: taskmgr,
		userusecase: usermgr,
	}
}



// GetTasks retrieves all tasks for the given user or admin.
func (controller *TaskController) GetTasks(c *gin.Context) {
	// Extract role and user ID from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	// Convert user ID from string to ObjectID
	ido, _ := primitive.ObjectIDFromHex(userid)

	// Fetch all tasks using the task usecase
	tasks, err := controller.taskusecase.GetAllTasks(c, role, ido)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tasks not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetTasksById retrieves a specific task by its ID.
func (controller *TaskController) GetTasksById(c *gin.Context) {
	id := c.Param("id") // Extract the task ID from the URL path

	// Extract role and user ID from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	// Fetch the task using the task usecase
	task, err := controller.taskusecase.GetTask(c, id, role, userid)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

// AddTask creates a new task for the user.
func (controller *TaskController) AddTask(c *gin.Context) {
	// Extract user ID from the context
	userid := c.GetString("userid")

	var task domain.Task

	// Bind the JSON request body to the task object
	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	// Convert user ID from string to ObjectID
	userido, _ := primitive.ObjectIDFromHex(userid)
	task.UserID = userido

	// Add the new task using the task usecase
	controller.taskusecase.AddTask(c, task)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task successfully created"})
}

// Updatetask updates an existing task identified by its ID.
func (controller *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id") // Extract the task ID from the URL path

	// Extract role and user ID from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	var updatedTask domain.Task

	// Bind the JSON request body to the updated task object
	err := c.ShouldBindJSON(&updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}

	// Convert user ID from string to ObjectID
	userido, _ := primitive.ObjectIDFromHex(userid)
	updatedTask.UserID = userido

	// Update the task using the task usecase
	erro := controller.taskusecase.UpdateTask(c, id, updatedTask, role)

	if erro != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task update failed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task successfully updated"})
}

// DeleteTask removes a task identified by its ID.
func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id") // Extract the task ID from the URL path

	// Extract role and user ID from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	// Delete the task using the task usecase
	erro := controller.taskusecase.DeleteTask(c, id, userid, role)
	if erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}


// RegisterUser handles the registration of a new user.
func (controller *TaskController) RegisterUser(c *gin.Context) {
	var user domain.User

	// Bind the JSON request body to the user object
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// Register the user using the user usecase
	req_status, err := controller.userusecase.RegisterUser(c, user)
	if err != nil {
		c.JSON(req_status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
}

// LoginUser handles user login and returns a token.
func (controller *TaskController) LoginUser(c *gin.Context) {
	var user domain.User

	// Bind the JSON request body to the user object
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	// Authenticate the user and get a token
	code, token, err := controller.userusecase.LoginUser(c, user)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully logged in", "token": token})
}

// DeleteUser removes a user identified by their ID (admin operation).
func (controller *TaskController) DeleteUser(c *gin.Context) {
	id := c.Param("id") // Extract the user ID from the URL path

	// Delete the user using the user usecase
	code, erro := controller.userusecase.DeleteUser(c, id)
	if erro == nil {
		c.IndentedJSON(code, gin.H{"message": "User successfully deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
