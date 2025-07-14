package auth

import (
	"advpractice/internal/user"
	"errors"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (service *AuthService) register(data *RegistrateRequest) (string, error) {
	existingUser, _ := service.UserRepository.GetByEmail(data.Email)
	if existingUser != nil {
		return "", errors.New(ErrUserExists)
	}

	user := &user.User{
		Name:     data.Name,
		Password: "", // Mock
		Email:    data.Email,
	}

	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
