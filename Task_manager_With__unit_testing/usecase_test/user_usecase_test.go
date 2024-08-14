package usecase_test

import (
	"context"
	"testing"
	Domain "task-manager/Domain"
	Usecase "task-manager/Usecases"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"errors"
	// "github.com/stretchr/testify/mock"
)

type UserUsecaseSuite struct {
	suite.Suite
	usecase    Domain.UserUsecase
	mockRepo   *MockUserRepository
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepository)
	suite.usecase = Usecase.NewUserUsecase(suite.mockRepo)
}

func (suite *UserUsecaseSuite) TestRegisterUser_Success() {
	user := Domain.User{Email: "test@example.com", Password: "password123"}

	suite.mockRepo.On("RegisterUser", user).Return(200, nil)

	status, err := suite.usecase.RegisterUser(context.Background(), user)
	suite.NoError(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestRegisterUser_Failure() {
	user := Domain.User{Email: "test@example.com", Password: "password123"}

	suite.mockRepo.On("RegisterUser", user).Return(0, errors.New("user already exists"))

	status, err := suite.usecase.RegisterUser(context.Background(), user)
	suite.Error(err)
	suite.Equal(0, status)
}

func (suite *UserUsecaseSuite) TestLoginUser_Success() {
	user := Domain.User{Email: "test@example.com", Password: "password123"}
	token := "fake-jwt-token"

	suite.mockRepo.On("LoginUser", user).Return(200, token, nil)

	status, resultToken, err := suite.usecase.LoginUser(context.Background(), user)
	suite.NoError(err)
	suite.Equal(200, status)
	suite.Equal(token, resultToken)
}

func (suite *UserUsecaseSuite) TestLoginUser_InvalidPassword() {
	user := Domain.User{Email: "test@example.com", Password: "wrongpassword"}

	suite.mockRepo.On("LoginUser", user).Return(401, "", errors.New("invalid credentials"))

	status, token, err := suite.usecase.LoginUser(context.Background(), user)
	suite.Error(err)
	suite.Equal(401, status)
	suite.Empty(token)
}

func (suite *UserUsecaseSuite) TestDeleteUser_Success() {
	userID := "some-user-id"

	suite.mockRepo.On("DeleteUser", userID).Return(200, nil)

	status, err := suite.usecase.DeleteUser(context.Background(), userID)
	suite.NoError(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestDeleteUser_Failure() {
	userID := "some-user-id"

	suite.mockRepo.On("DeleteUser", userID).Return(0, errors.New("user not found"))

	status, err := suite.usecase.DeleteUser(context.Background(), userID)
	suite.Error(err)
	suite.Equal(0, status)
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
