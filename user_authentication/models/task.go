package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    DueDate     time.Time          `bson:"due_date" json:"due_date"`
    Status      string             `bson:"status" json:"status"`
    OwnerID     primitive.ObjectID  `bson:"owner_id" json:"owner_id"`

}
