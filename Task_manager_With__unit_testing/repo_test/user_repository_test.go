package repo_test

import (
	Domain "task-manager/Domain"
	repo "task-manager/Repositories"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepositorySuite struct {
	suite.Suite
	repository *repo.MockUserRepository
}

func (suite *UserRepositorySuite) SetupSuite() {
	// Initialize the mock repository
	suite.repository = repo.NewMockUserRepository()
}

func (suite *UserRepositorySuite) TestRegisterUser() {
	// Create a new user
	user := Domain.User{
		Email:    "testuser@example.com",
		Password: "password123",
	}

	// Clear the repository for a fresh start
	suite.repository.Users = make(map[string]Domain.User)

	status, err := suite.repository.RegisterUser(user)
	suite.NoError(err, "no error should occur when registering a user")
	suite.Equal(200, status, "status code should be 200")

	// Verify the user was created
	registeredUser, exists := suite.repository.Users[user.Email]
	suite.True(exists, "user should be found in the repository")
	suite.Equal(user.Email, registeredUser.Email, "user email should match")
}

// func (suite *UserRepositorySuite) TestRegisterUser_AlreadyExists() {
// 	// Register a user
// 	user := Domain.User{
// 		Email:    "testuser@example.com",
// 		Password: "password123",
// 	}
// 	_, err := suite.repository.RegisterUser(user)
// 	suite.NoError(err, "no error should occur when registering a user")

// 	// Try to register the same user again
// 	status, err := suite.repository.RegisterUser(user)
// 	suite.Error(err, "error should occur when registering an existing user")
// 	suite.Equal(400, status, "status code should be 400")
// }

func (suite *UserRepositorySuite) TestLoginUser() {
	// Register a user
	user := Domain.User{
		Email:    "testuser@example.com",
		Password: "password123",
	}
	_, err := suite.repository.RegisterUser(user)
	suite.NoError(err, "no error should occur when registering a user")

	// Attempt to log in with the correct credentials
	loginUser := Domain.User{
		Email:    "testuser@example.com",
		Password: "password123",
	}
	status, token, err := suite.repository.LoginUser(loginUser)
	suite.NoError(err, "no error should occur when logging in")
	suite.Equal(200, status, "status code should be 200")
	suite.NotEmpty(token, "JWT token should not be empty")
}

func (suite *UserRepositorySuite) TestLoginUser_InvalidPassword() {
	// Register a user with a unique email to avoid conflicts
	user := Domain.User{
		Email:    "testuser_invalid@example.com",
		Password: "password123",
	}
	_, err := suite.repository.RegisterUser(user)
	suite.NoError(err, "no error should occur when registering a user")

	// Attempt to log in with an incorrect password
	loginUser := Domain.User{
		Email:    "testuser_invalid@example.com",
		Password: "wrongpassword",
	}
	status, token, err := suite.repository.LoginUser(loginUser)
	suite.Error(err, "error should occur when logging in with an incorrect password")
	suite.Equal(401, status, "status code should be 401")
	suite.Empty(token, "JWT token should be empty")
}

// Helper function to delete a user by email

func (suite *UserRepositorySuite) TestDeleteUser() {
	// Register a user
	user := Domain.User{
		Email:    "testuser@example.com",
		Password: "password123",
	}
	_, err := suite.repository.RegisterUser(user)
	suite.NoError(err, "no error should occur when registering a user")

	// Delete the user
	status, err := suite.repository.DeleteUser(user.ID.Hex())
	suite.NoError(err, "no error should occur when deleting a user")
	suite.Equal(200, status, "status code should be 200")

	// Verify the user was deleted
	_, exists := suite.repository.Users[user.Email]
	suite.False(exists, "user should not be found in the repository")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
