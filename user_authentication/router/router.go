package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"user_authentication/controllers"
	"user_authentication/data"
	"user_authentication/middleware"
)

func SetUpRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	userservice := data.NewUserService(client)
	taskService := data.NewTaskService(client)

	usercontroller := controllers.NewUserController(userservice)
	taskcontroller := controllers.NewTaskController(taskService)

	// User routes
	router.POST("/register", usercontroller.Register)
	router.POST("/login", usercontroller.Login)

	// Task routes
	router.POST("/tasks", middleware.AuthMiddleware(), taskcontroller.AddTask)
	router.GET("/tasks", middleware.AuthMiddleware(), taskcontroller.GetTasks)
	router.PUT("/tasks/:task_id", middleware.AuthMiddleware(), taskcontroller.UpdateTask) // Update task
	router.DELETE("/tasks/:task_id", middleware.AuthMiddleware(), taskcontroller.DeleteTask) // Delete task

	return router
}




























