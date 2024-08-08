package data

import (
	"context"
	"errors"
	"user_authentication/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(client *mongo.Client) *UserService {
	return &UserService{
		collection: client.Database("user_based_task_manager").Collection("users"),
	}
}

func (s *UserService) RegisterUser(user models.User) error {
	filter := bson.M{"username": user.Username}

	count, err := s.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already used")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	_, err = s.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) LoginUser(username, password string) (*models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := s.collection.FindOne(context.TODO(), filter).Decode(&user)
	
	
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil

}

// CountUsers counts the number of users matching a given ID
func (s *UserService) CountUsers(userID primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": userID}
	return s.collection.CountDocuments(context.TODO(), filter)
}

// GetUserRole retrieves the role of a user by their ID
func (s *UserService) GetUserRole(userID primitive.ObjectID) (string, error) {
	var user models.User
	err := s.collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}
