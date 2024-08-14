package repo_test

import (
    "testing"
    Domain "task-manager/Domain"
    repo "task-manager/Repositories"
    "github.com/stretchr/testify/suite"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepositorySuite struct {
    suite.Suite
    repository *repo.MockTaskRepository
}

func (suite *TaskRepositorySuite) SetupSuite() {
    // Initialize the mock repository
    suite.repository = repo.NewMockTaskRepository()
}

func (suite *TaskRepositorySuite) TestCreateTask() {
    // Create a new task
    task := Domain.Task{
        ID:          generateMockID(), // Implement a function to generate mock IDs
        Title:       "Test Task",
        Description: "This is a test task",
        Status:      "Pending",
    }

    err := suite.repository.AddTask(task)
    suite.NoError(err, "no error should occur when creating a task")

    // Verify the task was created
    createdTask, exists := suite.repository.Tasks[task.ID.Hex()]
    suite.True(exists, "task should be found in the repository")
    suite.Equal(task.Title, createdTask.Title, "task title should match")
}

func (suite *TaskRepositorySuite) TestGetTask() {
    // Create a new task
    task := Domain.Task{
        ID:          generateMockID(), // Implement a function to generate mock IDs
        Title:       "Test Task",
        Description: "This is a test task",
        Status:      "Pending",
    }
    suite.repository.AddTask(task)

    // Retrieve the task
    fetchedTask, err := suite.repository.GetTask(task.ID.Hex(), true, "")
    suite.NoError(err, "no error should occur when retrieving a task")
    suite.Equal(task.Title, fetchedTask.Title, "retrieved task title should match")
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
    // Create a new task
    task := Domain.Task{
        ID:          generateMockID(), // Implement a function to generate mock IDs
        Title:       "Test Task",
        Description: "This is a test task",
        Status:      "Pending",
    }
    suite.repository.AddTask(task)

    // Delete the task
    err := suite.repository.DeleteTask(task.ID.Hex())
    suite.NoError(err, "no error should occur when deleting a task")

    // Verify the task was deleted
    _, exists := suite.repository.Tasks[task.ID.Hex()]
    suite.False(exists, "task should not be found in the repository")
}

func generateMockID() primitive.ObjectID {
    // Implement a function to generate mock IDs
    return primitive.NewObjectID()
}

func TestTaskRepository(t *testing.T) {
    suite.Run(t, new(TaskRepositorySuite))
}
