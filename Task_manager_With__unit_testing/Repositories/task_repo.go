package repositories

import (
    "errors"
    "sync"
    Domain "task-manager/Domain"
    //"net/http"
)

type MockTaskRepository struct {
    mu    sync.Mutex
    Tasks map[string]Domain.Task
}

func NewMockTaskRepository() *MockTaskRepository {
    return &MockTaskRepository{
        Tasks: make(map[string]Domain.Task),
    }
}

func (repo *MockTaskRepository) AddTask(task Domain.Task) error {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    // Simulate task ID assignment
    if _, exists := repo.Tasks[task.ID.Hex()]; exists {
        return errors.New("task already exists")
    }

    repo.Tasks[task.ID.Hex()] = task
    return nil
}

func (repo *MockTaskRepository) GetTask(id string, _ bool, _ string) (Domain.Task, error) {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    task, exists := repo.Tasks[id]
    if !exists {
        return Domain.Task{}, errors.New("task not found")
    }

    return task, nil
}

func (repo *MockTaskRepository) DeleteTask(id string) error {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    if _, exists := repo.Tasks[id]; !exists {
        return errors.New("task not found")
    }

    delete(repo.Tasks, id)
    return nil
}
