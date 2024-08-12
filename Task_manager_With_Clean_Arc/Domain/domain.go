package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents the task data model.
type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	UserID      primitive.ObjectID `bson:"userid" json:"-"`
	Title       string             `bson:"title" json:"title,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	DueDate     time.Time          `bson:"due_date" json:"due_date,omitempty"`
	Status      string             `bson:"status" json:"status,omitempty"`
}

// TaskRepository defines the methods for interacting with the task data storage.
type TaskRepository interface {
	// GetAllTasks retrieves all tasks for a user or all users if admin.
	GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]Task, error)
	// GetTask retrieves a specific task by ID.
	GetTask(id string, isadmin bool, userid string) (Task, error)
	// AddTask creates a new task in the storage.
	AddTask(task Task) error
	// SetTask updates an existing task by ID.
	UpdateTask(id string, updatedTask Task, isadmin bool) error
	// DeleteTask removes a task by ID.
	DeleteTask(id string, userid string, isadmin bool) error
}

// TaskUsecase defines the business logic methods for tasks.
type TaskUsecase interface {
	// GetAllTasks retrieves all tasks for a user or all users if admin.
	GetAllTasks(c context.Context, isadmin bool, userid primitive.ObjectID) ([]Task, error)
	// GetTask retrieves a specific task by ID.
	GetTask(c context.Context, id string, isadmin bool, userid string) (Task, error)
	// AddTask creates a new task in the storage.
	AddTask(c context.Context, task Task) error
	// SetTask updates an existing task by ID.
	UpdateTask(c context.Context, id string, updatedTask Task, isadmin bool) error
	// DeleteTask removes a task by ID.
	DeleteTask(c context.Context, id string, userid string, isadmin bool) error
}

/*
=========== The Models and Interfaces for User =============
*/

// User represents the user data model.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	IsAdmin  bool               `json:"isadmin"`
}

// UserRepository defines the methods for interacting with the user data storage.
type UserRepository interface {
	// RegisterUserDb registers a new user in the database.
	RegisterUser(user User) (int, error)
	// LoginUserDb authenticates a user and returns a status code and token.
	LoginUser(user User) (int, string, error)
	// DeleteUser removes a user by ID.
	DeleteUser(id string) (int, error)
}

// UserUsecase defines the business logic methods for users.
type UserUsecase interface {
	// RegisterUserDb registers a new user in the database.
	RegisterUser(c context.Context, user User) (int, error)
	// LoginUserDb authenticates a user and returns a status code and token.
	LoginUser(c context.Context, user User) (int, string, error)
	// DeleteUser removes a user by ID.
	DeleteUser(c context.Context, id string) (int, error)
}
