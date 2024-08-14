package repositories

import (
    "errors"
    "sync"
    Domain "task-manager/Domain"
    "net/http"
)

type MockUserRepository struct {
    mu    sync.Mutex
    Users map[string]Domain.User
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        Users: make(map[string]Domain.User),
    }
}

func (repo *MockUserRepository) RegisterUser(user Domain.User) (int, error) {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    if _, exists := repo.Users[user.Email]; exists {
        return http.StatusBadRequest, errors.New("user already exists")
    }

    repo.Users[user.Email] = user
    return http.StatusOK, nil
}

func (repo *MockUserRepository) LoginUser(user Domain.User) (int, string, error) {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    storedUser, exists := repo.Users[user.Email]
    if !exists || storedUser.Password != user.Password {
        return http.StatusUnauthorized, "", errors.New("invalid email or password")
    }

    return http.StatusOK, "mock-jwt-token", nil
}

func (repo *MockUserRepository) DeleteUser(id string) (int, error) {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    for email, user := range repo.Users {
        if user.ID.Hex() == id {
            delete(repo.Users, email)
            return http.StatusOK, nil
        }
    }

    return http.StatusNotFound, errors.New("user not found")
}
