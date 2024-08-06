package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *data.TaskService
}

// NewTaskController creates a new instance of TaskController
func NewTaskController(ts *data.TaskService) *TaskController {
	return &TaskController{
		TaskService: ts,
	}
}

// GetAll_c handles GET requests to fetch all tasks
func (tc *TaskController) GetAll_c(c *gin.Context) {
	tasks, err := tc.TaskService.GetAll_d()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetById_c handles GET requests to fetch a task by ID
func (tc *TaskController) GetById_c(c *gin.Context) {
	id := c.Param("id")
	task, found, err := tc.TaskService.GetById_d(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// AddTask_c handles POST requests to add a new task
func (tc *TaskController) AddTask_c(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := tc.TaskService.AddTask_d(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, createdTask)
}

// Update_c handles PUT requests to update an existing task
func (tc *TaskController) Update_c(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := tc.TaskService.Update_d(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// Delete_c handles DELETE requests to delete a task
func (tc *TaskController) Delete_c(c *gin.Context) {
	id := c.Param("id")
	if err := tc.TaskService.Delete_d(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
