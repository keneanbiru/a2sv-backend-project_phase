package usecase_test

import (
	
	Domain "task-manager/Domain"
	
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) RegisterUser(user Domain.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) LoginUser(user Domain.User) (int, string, error) {
	args := m.Called(user)
	return args.Int(0), args.String(1), args.Error(2)
}

func (m *MockUserRepository) DeleteUser(id string) (int, error) {
	args := m.Called(id)
	return args.Int(0), args.Error(1)
}
