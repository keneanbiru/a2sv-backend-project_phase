package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

func SetRouter(ts *data.TaskService) *gin.Engine {
	r := gin.Default()
	taskController := controllers.NewTaskController(ts)

	api := r.Group("/api")
	{
		api.GET("/tasks", taskController.GetAll_c)
		api.GET("/tasks/:id", taskController.GetById_c)
		api.POST("/tasks", taskController.AddTask_c)
		api.PUT("/tasks/:id", taskController.Update_c)
		api.DELETE("/tasks/:id", taskController.Delete_c)
	}

	return r
}
