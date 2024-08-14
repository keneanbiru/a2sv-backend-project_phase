package usecase_test

import (
	
	Domain "task-manager/Domain"
	//Usecase "task-manager/Usecases"

	"github.com/stretchr/testify/mock"
	//"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTaskRepository is a mock implementation of the TaskRepository interface.
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]Domain.Task, error) {
	args := m.Called(isadmin, userid)
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTask(id string, isadmin bool, userid string) (Domain.Task, error) {
	args := m.Called(id, isadmin, userid)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) AddTask(task Domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) UpdateTask(id string, updatedTask Domain.Task, isadmin bool) error {
	args := m.Called(id, updatedTask, isadmin)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string, userid string, isadmin bool) error {
	args := m.Called(id, userid, isadmin)
	return args.Error(0)
}
