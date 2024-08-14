package usecases

import (
	"context"
	domain "task-manager/Domain"
)

// UserUsecase implements the UserUsecase interface.
type UserUsecase struct {
	userRepository domain.UserRepository
}

// NewUserUsecase creates a new instance of UserUsecase.
func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

// RegisterUser handles user registration by delegating to the repository.
func (useru *UserUsecase) RegisterUser(c context.Context, user domain.User) (int, error) {
	// Delegate the registration to the UserRepository
	return useru.userRepository.RegisterUser(user)
}

// LoginUser handles user login by delegating to the repository and returns a JWT token.
func (useru *UserUsecase) LoginUser(c context.Context, user domain.User) (int, string, error) {
	// Delegate the login to the UserRepository
	return useru.userRepository.LoginUser(user)
}

// DeleteUser handles the deletion of a user by delegating to the repository.
func (useru *UserUsecase) DeleteUser(c context.Context, id string) (int, error) {
	// Delegate the deletion to the UserRepository
	return useru.userRepository.DeleteUser(id)
}
