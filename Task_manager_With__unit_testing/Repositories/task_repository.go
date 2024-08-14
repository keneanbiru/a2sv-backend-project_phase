package repositories

import (
	"context"
	"errors"
	"log"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskRepository implements the TaskRepository interface for MongoDB.
type TaskRepository struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

// NewTaskRepository creates a new instance of TaskRepository.
func NewTaskRepository(mongoClient *mongo.Client) domain.TaskRepository {
	return &TaskRepository{
		Database:   mongoClient.Database("testing"),
		Collection: mongoClient.Database("testing").Collection("tasks"),
	}
}

// GetAllTasks retrieves all tasks for a user or all users if admin.
func (taskrepo *TaskRepository) GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]domain.Task, error) {
	var tasks []domain.Task

	filter := bson.M{}
	if !isadmin {
		filter = bson.M{"userid": userid}
	}

	cursor, err := taskrepo.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		if len(tasks) == 0 {
			return tasks, errors.New("no tasks found for the user")
		}
		return nil, err
	}
	return tasks, nil
}

// GetTask retrieves a specific task by ID.
func (taskrepo *TaskRepository) GetTask(id string, isadmin bool, userid string) (domain.Task, error) {
	var task domain.Task
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)

	filter := bson.M{"_id": ido}
	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	err := taskrepo.Collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return task, errors.New("task not found or you don't have permission to access it")
	}
	return task, nil
}

// AddTask creates a new task in the storage.
func (taskrepo *TaskRepository) AddTask(task domain.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := taskrepo.Collection.InsertOne(context.TODO(), task)

	if err != nil {
		return errors.New("failed to create task")
	}

	return nil
}

// SetTask updates an existing task by ID.
func (taskrepo *TaskRepository) UpdateTask(id string, updatedTask domain.Task, isadmin bool) error {
	ido, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": ido}
	if !isadmin {
		filter = bson.M{"_id": ido, "userid": updatedTask.UserID}
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
		},
	}

	result, err := taskrepo.Collection.UpdateOne(context.TODO(), filter, update)

	log.Println(result, err, updatedTask)

	if err != nil || result.MatchedCount == 0 {
		return errors.New("task not found or you lack the privilege to update it")
	}

	return nil
}

// DeleteTask removes a task by ID.
func (taskrepo *TaskRepository) DeleteTask(id string, userid string, isadmin bool) error {
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.M{"_id": ido}

	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	result, err := taskrepo.Collection.DeleteOne(context.TODO(), filter)

	if err != nil || result.DeletedCount == 0 {
		return errors.New("task not found or you don't have permission to delete it")
	}

	return nil
}
