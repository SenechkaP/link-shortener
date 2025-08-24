package auth

import (
	"advpractice/internal/user"
	"advpractice/pkg/di"
	"advpractice/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (service *AuthService) login(data *LoginRequest) (*uint, error) {
	user, err := service.UserRepository.GetByEmail(data.Email)
	if err != nil {
		return nil, errors.New(ErrNonExistentEmail)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, errors.New(ErrWrongPassword)
	}

	return &user.ID, nil
}

func (service *AuthService) register(data *RegistrateRequest) (*uint, error) {
	existingUser, _ := service.UserRepository.GetByEmail(data.Email)
	if existingUser != nil {
		return nil, errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &user.User{
		Name:     data.Name,
		Password: string(hashedPassword),
		Email:    data.Email,
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (service *AuthService) getJWT(data jwt.JWTData, secret string) (string, error) {
	j := jwt.NewJWT(secret)
	token, err := j.Create(data)
	if err != nil {
		return "", errors.New(ErrJWT)
	}

	return token, nil
}
