package repositories

import (
	"context"
	"errors"
	"log"
	"net/http"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository implements the UserRepository interface for MongoDB.
type UserRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(mongoClient *mongo.Client) domain.UserRepository {
	return &UserRepository{
		database:   mongoClient.Database("task-user"),
		collection: mongoClient.Database("task-user").Collection("users"), // Changed collection name to "users"
	}
}

// RegisterUserDb registers a new user in the database.
func (userepo *UserRepository) RegisterUser(user domain.User) (int, error) {
	collection := userepo.collection

	// Check if a user with the same email already exists
	ere := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Err()
	if ere == nil {
		return http.StatusBadRequest, errors.New("user already exists with the same email")
	}

	// Hash the password before storing it
	hashedPassword, err := infrastructure.PasswordHasher(user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	user.Password = hashedPassword

	// Generate a new ObjectID for the user
	user.ID = primitive.NewObjectID()

	// Insert the new user into the collection
	_, erro := collection.InsertOne(context.TODO(), user)
	if erro != nil {
		return http.StatusBadRequest, erro
	}

	return http.StatusOK, nil
}

// LoginUserDb authenticates a user and returns a JWT token.
func (userepo *UserRepository) LoginUser(user domain.User) (int, string, error) {
	collection := userepo.collection

	var existingUser domain.User

	// Retrieve the user with the provided email
	collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	log.Println(existingUser, user)

	// Check if the provided password matches the stored hashed password
	if !infrastructure.PasswordComparator(existingUser.Password, user.Password) {
		return http.StatusUnauthorized, "", errors.New("invalid email or password")
	}

	// Generate a JWT token for the authenticated user
	jwtToken, err := infrastructure.TokenGenerator(existingUser.ID, existingUser.Email, existingUser.IsAdmin)
	if err != nil {
		return http.StatusInternalServerError, "", errors.New("internal server error")
	}

	return http.StatusOK, jwtToken, nil
}

// DeleteUser removes a user by ID.
func (userepo *UserRepository) DeleteUser(id string) (int, error) {
	collection := userepo.collection

	ido, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ido}

	// Delete the user from the collection
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil || result.DeletedCount == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	return http.StatusOK, nil
}
