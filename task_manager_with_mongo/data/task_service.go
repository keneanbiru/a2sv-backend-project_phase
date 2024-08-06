package data

import (
	"context"
	"fmt"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

// NewTaskService creates a new TaskService with a MongoDB collection
func NewTaskService(client *mongo.Client) *TaskService {
	collection := client.Database("taskdb").Collection("tasks")
	return &TaskService{
		collection: collection,
	}
}

// GetAll_d retrieves all tasks from the database
func (ts *TaskService) GetAll_d() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := ts.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask_d adds a new task to the database
func (ts *TaskService) AddTask_d(task models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	task.DueDate = time.Now() // Set DueDate to current time for example
	_, err := ts.collection.InsertOne(context.Background(), task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// GetById_d retrieves a task by ID from the database
func (ts *TaskService) GetById_d(id string) (models.Task, bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, false, fmt.Errorf("invalid ID format")
	}

	var task models.Task
	err = ts.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, false, nil
		}
		return models.Task{}, false, err
	}

	return task, true, nil
}

// Update_d updates an existing task in the database
func (ts *TaskService) Update_d(id string, updated models.Task) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, fmt.Errorf("invalid ID format")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
			"status":      updated.Status,
		},
	}

	result := ts.collection.FindOneAndUpdate(context.Background(), filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return models.Task{}, fmt.Errorf("task not found")
		}
		return models.Task{}, result.Err()
	}

	// Fetch the updated document
	var updatedTask models.Task
	err = result.Decode(&updatedTask)
	if err != nil {
		return models.Task{}, err
	}

	return updatedTask, nil
}

// Delete_d deletes a task by ID from the database
func (ts *TaskService) Delete_d(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}

	_, err = ts.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
