package usecase_test

import (
	"context"
	"testing"
	Domain "task-manager/Domain"
	Usecase "task-manager/Usecases"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseSuite struct {
	suite.Suite
	usecase    Domain.TaskUsecase
	mockRepo   *MockTaskRepository
}

func (suite *TaskUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockTaskRepository)
	suite.usecase = Usecase.NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskUsecaseSuite) TestGetAllTasks_Success() {
	userid := primitive.NewObjectID()
	expectedTasks := []Domain.Task{{Title: "Task 1"}, {Title: "Task 2"}}

	suite.mockRepo.On("GetAllTasks", false, userid).Return(expectedTasks, nil)

	tasks, err := suite.usecase.GetAllTasks(context.Background(), false, userid)
	suite.NoError(err)
	suite.Equal(expectedTasks, tasks)
}

func (suite *TaskUsecaseSuite) TestGetTask_Success() {
	taskID := "task-id"
	userid := "user-id"
	expectedTask := Domain.Task{Title: "Task 1"}

	suite.mockRepo.On("GetTask", taskID, false, userid).Return(expectedTask, nil)

	task, err := suite.usecase.GetTask(context.Background(), taskID, false, userid)
	suite.NoError(err)
	suite.Equal(expectedTask, task)
}

func (suite *TaskUsecaseSuite) TestAddTask_Success() {
	newTask := Domain.Task{Title: "New Task"}

	suite.mockRepo.On("AddTask", newTask).Return(nil)

	err := suite.usecase.AddTask(context.Background(), newTask)
	suite.NoError(err)
}

func (suite *TaskUsecaseSuite) TestUpdateTask_Success() {
	taskID := "task-id"
	updatedTask := Domain.Task{Title: "Updated Task"}

	suite.mockRepo.On("UpdateTask", taskID, updatedTask, false).Return(nil)

	err := suite.usecase.UpdateTask(context.Background(), taskID, updatedTask, false)
	suite.NoError(err)
}

func (suite *TaskUsecaseSuite) TestDeleteTask_Success() {
	taskID := "task-id"
	userid := "user-id"

	suite.mockRepo.On("DeleteTask", taskID, userid, false).Return(nil)

	err := suite.usecase.DeleteTask(context.Background(), taskID, userid, false)
	suite.NoError(err)
}

func TestTaskUsecase(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
