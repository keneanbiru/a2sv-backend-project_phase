package controllers

import (
	"net/http"
	// "strconv"
	"task_manager_api/data"
	"task_manager_api/models"
	// "task_manager_api/models"
	"github.com/gin-gonic/gin"
)

type TaskController struct{
	TaskService *data.Taskservice
}

func NewTaskController (ts *data.Taskservice) *TaskController{
return &TaskController{
TaskService: ts,
}
}

func (ts *TaskController) GetAll_c(c *gin.Context){
tasks := ts.TaskService.GetAll_d()
// if ts.TaskService.Lastid != 1{
c.IndentedJSON(http.StatusOK,tasks)

//c.IndentedJSON(http.StatusNoContent,tasks)
}


func (ts *TaskController) GetById_c (c *gin.Context){
    id := c.Param("id")
    task,ex := ts.TaskService.GetById_d(id)
    if !ex {
        c.JSON(http.StatusOK, gin.H{"error": "Task not found"})
        return

    }
    c.IndentedJSON(http.StatusOK,task)

}

func (ts *TaskController) AddTask_c (c *gin.Context) {
    var task models.Task
    if err :=c.BindJSON(&task) ; err!= nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return 
    }
    
    CreatedTask := ts.TaskService.AddTask_d(task)
    c.JSON(http.StatusOK,CreatedTask )
}

func (ts *TaskController) Update_c (c *gin.Context){
    id := c.Param("id")
    var task models.Task
    if err :=c.BindJSON(&task) ; err!= nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return 
    }
    updatedtask,err  := ts.TaskService.Update_d(id,task)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK,updatedtask)
    
}

func (ts *TaskController) Delete_c (c *gin.Context){
    id := c.Param("id")
    err := ts.TaskService.Delete_d(id)

    if err != nil{
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return 
    }
    c.IndentedJSON(http.StatusOK,nil)


}