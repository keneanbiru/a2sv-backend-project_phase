package controller_test

import (
	//"task-manager/Domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockTaskController is a mock implementation of TaskController
type MockTaskController struct {
	mock.Mock
}

func (m *MockTaskController) GetTasks(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) GetTasksById(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) AddTask(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) UpdateTask(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) DeleteTask(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) RegisterUser(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) LoginUser(c *gin.Context) {
	m.Called(c)
}

func (m *MockTaskController) DeleteUser(c *gin.Context) {
	m.Called(c)
}

// MockUserController is a mock implementation of UserController
type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) Register(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserController) Login(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserController) GetUsers(c *gin.Context) {
	m.Called(c)
}
