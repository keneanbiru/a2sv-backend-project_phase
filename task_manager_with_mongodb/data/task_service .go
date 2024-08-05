package data

import (
	// "strconv"
	"task_manager_api/models"
	"fmt"
	// "error"
)

type Taskservice struct{
	Tasks  map[string]models.Task
	Lastid int
}

func Newtaskservice () *Taskservice{
	return &Taskservice{
		Tasks :  make(map[string]models.Task),
		Lastid : 1,

}
}

func (ts *Taskservice) GetAll_d () []models.Task{
	var task []models.Task
	for _,tasks := range ts.Tasks{
		task = append(task,tasks)
	}
	return task
}

func (ts *Taskservice) AddTask_d(task models.Task) models.Task {
	task.ID = fmt.Sprintf("task-%d", ts.Lastid) // Generate string ID
	ts.Tasks [task.ID] = task
	ts.Lastid +=1
	return task
}


func (ts *Taskservice) GetById_d (id string) ( models.Task,bool){
task,ex := ts.Tasks[id]
return task,ex
}

func (ts *Taskservice) Update_d  (id string,updated models.Task) (models.Task,error){
	task,ex := ts.Tasks[id]
	if !ex {
		return models.Task{}, fmt.Errorf("task not found")
	}
	task.ID = updated.ID
	task.Description = updated.Description
	task.Title = updated.Description
	task.DueDate = updated.DueDate
	task.Status = updated.Status
	ts.Tasks[id] = task
	return task,nil

}

func  (ts *Taskservice) Delete_d(id string)  error{
	_,ex :=  ts.Tasks[id]
	if !ex{
		return fmt.Errorf("task not found")
	}
	delete(ts.Tasks,id)
	return nil
} 