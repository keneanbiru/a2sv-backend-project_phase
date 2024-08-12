package usecase

import (
	"context"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// taskUsecase implements the TaskUsecase interface.
type taskUsecase struct {
	taskRepository domain.TaskRepository
}

// NewTaskUsecase creates a new instance of taskUsecase.
func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

// GetAllTasks retrieves all tasks based on user role and ID.
func (usecase *taskUsecase) GetAllTasks(c context.Context, isadmin bool, userid primitive.ObjectID) ([]domain.Task, error) {
	// Delegate the retrieval to the TaskRepository
	return usecase.taskRepository.GetAllTasks(isadmin, userid)
}

// GetTask retrieves a specific task by its ID and checks permissions based on role and user ID.
func (usecase *taskUsecase) GetTask(c context.Context, id string, isadmin bool, userid string) (domain.Task, error) {
	// Delegate the retrieval to the TaskRepository
	return usecase.taskRepository.GetTask(id, isadmin, userid)
}

// AddTask creates a new task.
func (usecase *taskUsecase) AddTask(c context.Context, task domain.Task) error {
	// Delegate the addition to the TaskRepository
	return usecase.taskRepository.AddTask(task)
}

// UpdateTask updates an existing task by its ID.
func (usecase *taskUsecase) UpdateTask(c context.Context, id string, updatedTask domain.Task, isadmin bool) error {
	// Delegate the update to the TaskRepository
	return usecase.taskRepository.UpdateTask(id, updatedTask, isadmin)
}

// DeleteTask removes a task by its ID, checking permissions based on role and user ID.
func (usecase *taskUsecase) DeleteTask(c context.Context, id string, userid string, isadmin bool) error {
	// Delegate the deletion to the TaskRepository
	return usecase.taskRepository.DeleteTask(id, userid, isadmin)
}
