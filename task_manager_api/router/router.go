package router

import (
	"task_manager_api/controllers"
	"task_manager_api/data"

	"github.com/gin-gonic/gin"
)

func SetRouter(ts *data.Taskservice) *gin.Engine {
	r := gin.Default()
	taskcontroller := controllers.NewTaskController(ts)
	api := r.Group("/api")
	{
		api.GET("/tasks", taskcontroller.GetAll_c)
		api.GET("/tasks/:id", taskcontroller.GetById_c)
		api.POST("/tasks", taskcontroller.AddTask_c)
		api.PUT("/tasks/:id", taskcontroller.Update_c)
		api.DELETE("/tasks/:id", taskcontroller.Delete_c)
	}
	return r
}
