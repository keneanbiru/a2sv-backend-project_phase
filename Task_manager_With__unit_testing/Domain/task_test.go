package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTask_Initialization(t *testing.T) {
	// Set up test data
	id := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	dueDate := time.Now().Add(24 * time.Hour)

	// Initialize the entity
	task := Task{
		ID:          id,
		UserID:      userID,
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     dueDate,
		Status:      "pending",
	}

	// Verify that the entity fields are set correctly
	assert.Equal(t, id, task.ID)
	assert.Equal(t, userID, task.UserID)
	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "This is a test task", task.Description)
	assert.Equal(t, dueDate, task.DueDate)
	assert.Equal(t, "pending", task.Status)
}

func TestTask_JSONSerialization(t *testing.T) {
	// Initialize the entity
	id := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	task := Task{
		ID:          id,
		UserID:      userID,
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(task)
	assert.NoError(t, err)

	// Deserialize back to a Task
	var deserializedTask Task
	err = json.Unmarshal(jsonData, &deserializedTask)
	assert.NoError(t, err)

	// Ensure the deserialized entity matches the original
	assert.Equal(t, task.ID, deserializedTask.ID)
	assert.Equal(t, task.UserID, deserializedTask.UserID)
	assert.Equal(t, task.Title, deserializedTask.Title)
	assert.Equal(t, task.Description, deserializedTask.Description)
	assert.True(t, task.DueDate.Equal(deserializedTask.DueDate), "DueDate should be equal")
	assert.Equal(t, task.Status, deserializedTask.Status)
}
