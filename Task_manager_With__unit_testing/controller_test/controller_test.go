package controller_test

import (
	//"task-manager/Domain"
	//"task-manager/"
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserController_Register(t *testing.T) {
	mockController := new(MockUserController)
	mockController.On("Register", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":"testuser@example.com","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockController.Register(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestUserController_Login(t *testing.T) {
	mockController := new(MockUserController)
	mockController.On("Login", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", strings.NewReader(`{"email":"testuser@example.com","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockController.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestUserController_GetUsers(t *testing.T) {
	mockController := new(MockUserController)
	mockController.On("GetUsers", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/users", nil)

	mockController.GetUsers(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_GetTasks(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("GetTasks", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/tasks", nil)

	mockController.GetTasks(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_GetTasksById(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("GetTasksById", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/tasks/1", nil)

	mockController.GetTasksById(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_AddTask(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("AddTask", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"title":"New Task"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockController.AddTask(c)

	assert.Equal(t, http.StatusOK, w.Code) // Expecting 201 Created
	mockController.AssertExpectations(t)
}


func TestTaskController_UpdateTask(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("UpdateTask", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/tasks/1", strings.NewReader(`{"title":"Updated Task"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockController.UpdateTask(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_DeleteTask(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("DeleteTask", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("DELETE", "/tasks/1", nil)

	mockController.DeleteTask(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_RegisterUser(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("RegisterUser", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":"testuser@example.com","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockController.RegisterUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_LoginUser(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("LoginUser", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", strings.NewReader(`{"email":"testuser@example.com","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockController.LoginUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}

func TestTaskController_DeleteUser(t *testing.T) {
	mockController := new(MockTaskController)
	mockController.On("DeleteUser", mock.Anything).Return()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("DELETE", "/users/1", nil)

	mockController.DeleteUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockController.AssertExpectations(t)
}
