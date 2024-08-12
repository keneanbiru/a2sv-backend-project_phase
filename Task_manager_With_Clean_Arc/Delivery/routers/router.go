package routers

import (
	"task-manager/Delivery/controllers"
	infrastructure "task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes and configures the routes for the application.
func SetupRoutes(gino *gin.Engine, taskmgr *controllers.TaskController) {

	// Public routes
	// Route for user registration
	gino.POST("/register", taskmgr.RegisterUser)
	// Route for user login
	gino.POST("/login", taskmgr.LoginUser)

	// Protected routes group
	auth := gino.Group("/")
	// Apply authentication middleware to protected routes
	auth.Use(infrastructure.AuthMiddleware)
	{
		auth.DELETE("/tasks/:id", taskmgr.DeleteTask)
		// Route to get all tasks
		auth.GET("/tasks/:id", taskmgr.GetTasksById)
		// Route to update a task by ID		
		auth.PUT("/tasks/:id", taskmgr.UpdateTask)
		// Route to create a new task
		auth.POST("/tasks", taskmgr.AddTask)
		// Route to delete a task by ID
		auth.GET("/tasks", taskmgr.GetTasks)
		// Route to get a specific task by ID
		

		// Admin-specific endpoint group
		admin := auth.Group("/")
		// Apply admin middleware to admin-specific routes
		admin.Use(infrastructure.AdminMiddleware)
		{
			// Route to delete a user by ID (admin operation)
			admin.DELETE("/users/:id", taskmgr.DeleteUser)
		}
	}
}
