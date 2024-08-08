package data

import (
	"context"
	"errors"
	"time"
	"user_authentication/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection  *mongo.Collection
	userService *UserService
}

func NewTaskService(client *mongo.Client) *TaskService {
	return &TaskService{
		collection: client.Database("user_based_task_manager").Collection("tasks"),
	}
}

func (ts *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	// Validate the OwnerID
	if task.OwnerID == primitive.NilObjectID {
		return nil, errors.New("invalid Owner ID")
	}

	// Ensure the task ID is set
	if task.ID == primitive.NilObjectID {
		task.ID = primitive.NewObjectID()
	}

	// Set a timeout context for the insertion operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Perform the insertion using the collection's InsertOne method
	result, err := ts.collection.InsertOne(ctx, task)
	if err != nil {
		// Handle different MongoDB errors if needed
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New("task with this ID already exists")
		}
		// Additional error handling for different scenarios
		if mongo.IsNetworkError(err) {
			return nil, errors.New("network error occurred while creating task: " + err.Error())
		}
		if mongo.IsTimeout(err) {
			return nil, errors.New("operation timed out while creating task: " + err.Error())
		}
		return nil, errors.New("failed to create task: " + err.Error())
	}

	return result, nil
}

func (ts *TaskService) Gettasks(userid primitive.ObjectID, role string) ([]models.Task, error) {
	// Filter all the tasks from any user if role is admin but filter only user-specific tasks if it is not admin
	var filter bson.M
	if role == "admin" {
		filter = bson.M{}
	} else {
		filter = bson.M{"owner_id": userid}
	}

	cursor, err := ts.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (ts *TaskService) UpdateTask(id primitive.ObjectID, updatedTask models.Task) (*mongo.UpdateResult, error) {
	// Set a timeout context for the update operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter for finding the task to update
	filter := bson.M{"_id": id}

	// Define the update fields
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
			"owner_id":    updatedTask.OwnerID,
		},
	}

	// Perform the update operation using the collection's UpdateOne method
	result, err := ts.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ts *TaskService) DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	// Set a timeout context for the deletion operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter for finding the task to delete
	filter := bson.M{"_id": id}

	// Perform the delete operation using the collection's DeleteOne method
	result, err := ts.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TaskService) UserExists(userID primitive.ObjectID) (bool, error) {
		count, err := s.userService.CountUsers(userID)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}








































// package data

// import (
// 	"context"
// 	"errors"
// 	"time"
// 	"user_authentication/models"

// 	// "go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type TaskService struct {
// 	collection  *mongo.Collection
// 	userService *UserService
// }

// func NewTaskService(client *mongo.Client) *TaskService {
// 	return &TaskService{
// 		collection: client.Database("task_manager").Collection("tasks"),
// 	}
// }


// func (ts *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
// 	// Validate the OwnerID
// 	if task.OwnerID == primitive.NilObjectID {
// 		return nil, errors.New("invalid Owner ID")
// 	}

// 	// Ensure the task ID is set
// 	if task.ID == primitive.NilObjectID {
// 		task.ID = primitive.NewObjectID()
// 	}

// 	// Set a timeout context for the insertion operation
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Perform the insertion using the collection's InsertOne method
// 	result, err := ts.collection.InsertOne(ctx, task)
// 	if err != nil {
// 		// Handle different MongoDB errors if needed
// 		if mongo.IsDuplicateKeyError(err) {
// 			return nil, errors.New("task with this ID already exists")
// 		}
// 		// Additional error handling for different scenarios
// 		if mongo.IsNetworkError(err) {
// 			return nil, errors.New("network error occurred while creating task: " + err.Error())
// 		}
// 		if mongo.IsTimeout(err) {
// 			return nil, errors.New("operation timed out while creating task: " + err.Error())
// 		}
// 		return nil, errors.New("failed to create task: " + err.Error())
// 	}

// 	return result, nil
// }


// func (ts *TaskService) Gettasks(userid primitive.ObjectID, role string) ([]models.Task, error) {
// 	//filter all the tasks from any user if role is admin but filter only user specific tasks if it is not admin
// 	var filter bson.M
// 	if role == "admin" {
// 		filter = bson.M{}
// 	} else {
// 		filter = bson.M{"owner_id": userid}
// 	}

// 	cursor, err := ts.collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	var tasks []models.Task
// 	for cursor.Next(context.TODO()) {
// 		var task models.Task
// 		if err := cursor.Decode(&task); err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, task)
// 	}

// 	return tasks, nil
// }
// func (s *TaskService) UserExists(userID primitive.ObjectID) (bool, error) {
// 	count, err := s.userService.CountUsers(userID)
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }
