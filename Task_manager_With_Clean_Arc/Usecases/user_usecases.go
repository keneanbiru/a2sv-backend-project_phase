package usecase

import (
	"context"
	domain "task-manager/Domain"
)

// userUsecase implements the UserUsecase interface.
type userUsecase struct {
	userRepository domain.UserRepository
}

// NewUserUsecase creates a new instance of userUsecase.
func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

// RegisterUser handles user registration by delegating to the repository.
func (useru *userUsecase) RegisterUser(c context.Context, user domain.User) (int, error) {
	// Delegate the registration to the UserRepository
	return useru.userRepository.RegisterUser(user)
}

// LoginUser handles user login by delegating to the repository and returns a JWT token.
func (useru *userUsecase) LoginUser(c context.Context, user domain.User) (int, string, error) {
	// Delegate the login to the UserRepository
	return useru.userRepository.LoginUser(user)
}

// DeleteUser handles the deletion of a user by delegating to the repository.
func (useru *userUsecase) DeleteUser(c context.Context, id string) (int, error) {
	// Delegate the deletion to the UserRepository
	return useru.userRepository.DeleteUser(id)
}
